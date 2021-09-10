package janitor

import (
	"context"
	"fmt"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/gobwas/glob"
	"github.com/hashicorp/go-multierror"
	lru "github.com/hashicorp/golang-lru"
	"github.com/inconshreveable/log15"

	"github.com/sourcegraph/sourcegraph/enterprise/internal/codeintel/gitserver"
	"github.com/sourcegraph/sourcegraph/enterprise/internal/codeintel/stores/dbstore"
	"github.com/sourcegraph/sourcegraph/internal/goroutine"
	"github.com/sourcegraph/sourcegraph/internal/timeutil"
)

type uploadExpirer struct {
	dbStore                DBStore
	gitserverClient        GitserverClient
	metrics                *metrics
	repositoryProcessDelay time.Duration
	repositoryBatchSize    int
	uploadProcessDelay     time.Duration
	uploadBatchSize        int
	commitBatchSize        int
	cacheMaxKeys           int
}

var _ goroutine.Handler = &uploadExpirer{}
var _ goroutine.ErrorHandler = &uploadExpirer{}

// NewUploadExpirer returns a background routine that periodically compares the age of upload records
// against the age of uploads protected by global and repository specific data retention policies.
//
// Uploads that are older than the protected retention age are marked as expired. Expired records with
// no dependents will be removed by the expiredUploadDeleter.
func NewUploadExpirer(
	dbStore DBStore,
	gitserverClient GitserverClient,
	repositoryProcessDelay time.Duration,
	repositoryBatchSize int,
	uploadProcessDelay time.Duration,
	uploadBatchSize int,
	commitBatchSize int,
	cacheMaxKeys int,
	interval time.Duration,
	metrics *metrics,
) goroutine.BackgroundRoutine {
	return goroutine.NewPeriodicGoroutine(context.Background(), interval, &uploadExpirer{
		dbStore:                dbStore,
		gitserverClient:        gitserverClient,
		metrics:                metrics,
		repositoryProcessDelay: repositoryProcessDelay,
		repositoryBatchSize:    repositoryBatchSize,
		uploadProcessDelay:     uploadProcessDelay,
		uploadBatchSize:        uploadBatchSize,
		commitBatchSize:        commitBatchSize,
		cacheMaxKeys:           cacheMaxKeys,
	})
}

func (e *uploadExpirer) Handle(ctx context.Context) (err error) {
	// Get the batch of repositories that we'll handle in this invocation of the periodic goroutine. This
	// set should contain repositories that have yet to be updated, or that have been updated least recently.
	// This allows us to update every repository reliably, even if it takes a long time to process through
	// the backlog.
	lastUpdatedAtByRepository, err := e.dbStore.SelectRepositoriesForRetentionScan(ctx, e.repositoryProcessDelay, e.repositoryBatchSize)
	if err != nil {
		return err
	}
	if len(lastUpdatedAtByRepository) == 0 {
		// All repositories updated recently enough
		return nil
	}

	// Retrieve the set of global configuration policies that affect data retention. These policies are
	// applied to all repositories.
	globalPolicies, err := e.dbStore.GetConfigurationPolicies(ctx, dbstore.GetConfigurationPoliciesOptions{
		ForDataRetention: true,
	})
	if err != nil {
		return err
	}

	for repositoryID, repositoryLastUpdatedAt := range lastUpdatedAtByRepository {
		if repositoryErr := e.handleRepository(ctx, repositoryID, repositoryLastUpdatedAt, globalPolicies); repositoryErr != nil {
			if err == nil {
				err = repositoryErr
			} else {
				err = multierror.Append(err, repositoryErr)
			}
		}
	}

	// TODO - add metrics for items processed
	return nil
}

func (e *uploadExpirer) HandleError(err error) {
	e.metrics.numErrors.Inc()
	log15.Error("Failed to expire old codeintel records", "error", err)
}

func (e *uploadExpirer) handleRepository(ctx context.Context, repositoryID int, repositoryLastUpdatedAt *time.Time, globalPolicies []dbstore.ConfigurationPolicy) error {
	// Retrieve the set of configuration policies that affect data retention. These policies are applied
	// only to this repository.
	repositoryPolicies, err := e.dbStore.GetConfigurationPolicies(ctx, dbstore.GetConfigurationPoliciesOptions{
		RepositoryID:     repositoryID,
		ForDataRetention: true,
	})
	if err != nil {
		return err
	}

	// Combine global and repository-specific policies
	policies := append(globalPolicies, repositoryPolicies...)

	// Construct a map from policy pattern to a compiled glob object used to match to commits, branch names,
	// and tag names. If there are multiple policies with the same pattern, the pattern is compiled only once.
	patterns := make(map[string]glob.Glob, len(policies))

	for _, policy := range policies {
		if _, ok := patterns[policy.Pattern]; ok {
			continue
		}

		pattern, err := glob.Compile(policy.Pattern)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to compile glob pattern `%s` in configuration policy %d", policy.Pattern, policy.ID))
		}

		patterns[policy.Pattern] = pattern
	}

	// Get a list of relevant branch and tag heads of this repository
	refDescriptions, err := e.gitserverClient.RefDescriptions(ctx, repositoryID)
	if err != nil {
		return err
	}

	// Create a cache structure shared by the routine that processes each upload. An upload can be
	// visible from many commits at once, so it is likely that the same commit is re-processed many
	// times. This cache prevents us from making redundant gitserver requests, and from wasting
	// compute time iterating through the same data already in memory.
	repositoryCache, err := newRepositoryCache(e.cacheMaxKeys)
	if err != nil {
		return err
	}

	// Mark the time after which all unprocessed uploads for this repository will not be touched.
	// This timestamp field is used as a rate limiting device so we do not busy-loop over the same
	// protected records in the background.
	//
	// This value should be assigned OUTSIDE of the following loop to prevent the case where the
	// upload process delay is shorter than the time it takes to process one batch of uploads. This
	// is obviously a mis-configuration, but one we can make a bit less catastrophic by not updating
	// this value in the loop.
	lastRetentionScanBefore := time.Now().Add(-e.uploadProcessDelay)

	for {
		// Each record pulled back by this query will either have its expired flag or its last
		// retention scan timestamp updated by the following handleUploads call. This guarantees
		// that the loop will terminate naturally after the entire set of candidate uploads have
		// been seen and updated with a time necessarily greater than lastRetentionScanBefore.

		uploads, _, err := e.dbStore.GetUploads(ctx, dbstore.GetUploadsOptions{
			State:                   "completed",
			RepositoryID:            repositoryID,
			OldestFirst:             true,
			Limit:                   e.uploadBatchSize,
			LastRetentionScanBefore: &lastRetentionScanBefore,
			UploadedBefore:          repositoryLastUpdatedAt,
		})
		if err != nil {
			return err
		}
		if len(uploads) == 0 {
			break
		}

		if err := e.handleUploads(ctx, policies, patterns, refDescriptions, repositoryCache, uploads); err != nil {
			return err
		}
	}

	return nil
}

func (e *uploadExpirer) handleUploads(
	ctx context.Context,
	policies []dbstore.ConfigurationPolicy,
	patterns map[string]glob.Glob,
	refDescriptions map[string][]gitserver.RefDescription,
	repositoryCache *repositoryCache,
	uploads []dbstore.Upload,
) (err error) {
	// Categorize each upload as protected or expired
	protectedUploadIDs := make([]int, 0, len(uploads))
	expiredUploadIDs := make([]int, 0, len(uploads))

	for _, upload := range uploads {
		protected, checkErr := e.isUploadProtectedByPolicy(
			ctx,
			policies,
			patterns,
			refDescriptions,
			repositoryCache,
			upload,
		)
		if checkErr != nil {
			if err == nil {
				err = checkErr
			} else {
				err = multierror.Append(err, checkErr)
			}

			// Collect errors but not prevent other commits from being successfully
			// processed. We'll leave the ones that fail here alone to be re-checked
			// the next time records for this repository are scanned.
			continue
		}

		if protected {
			protectedUploadIDs = append(protectedUploadIDs, upload.ID)
		} else {
			expiredUploadIDs = append(expiredUploadIDs, upload.ID)
		}
	}

	// Update the last data retention scan timestamp on the upload records with the given
	// protected identifiers (so that we do not re-select the same uploads on the next batch)
	// and sets the expired field on the upload records with the given expired identifiers
	// (so that the expiredUploadDeleter process can remove then once unreferenced).

	if updateErr := e.dbStore.UpdateUploadRetention(ctx, protectedUploadIDs, expiredUploadIDs); updateErr != nil {
		updateErr := errors.Wrap(err, "dbstore.UpdateUploadRetention")

		if err == nil {
			err = updateErr
		} else {
			err = multierror.Append(err, updateErr)
		}
	}

	return err
}

func (e *uploadExpirer) isUploadProtectedByPolicy(
	ctx context.Context,
	policies []dbstore.ConfigurationPolicy,
	patterns map[string]glob.Glob,
	refDescriptions map[string][]gitserver.RefDescription,
	repositoryCache *repositoryCache,
	upload dbstore.Upload,
) (bool, error) {
	// Determine the set of policies that will apply to the fast path. This excludes any policies that
	// do not cover the time of the upload. Any policy removed here can not protect the given upload.
	//
	// Also note that on each invocation of this method for the same repository, this set of policies
	// can only increase as we process uploads in descending age.
	fastPathPolicies := filterPolicies(policies, func(policy dbstore.ConfigurationPolicy) bool {
		return policyCoversUpload(policy, upload)
	})

	// Determine the set of policies that will apply to the slow path. This excludes any policies that
	// were not covered in the fast path, as well as any policies that do not cover intermediate commits
	// of a branch. The fast path, which runs before the slow path, will have already checked whether
	// these policies protect the upload as the tip of a branch - nwo we need to compare policies against
	// commits within a branch.
	slowPathPolicies := filterPolicies(fastPathPolicies, func(policy dbstore.ConfigurationPolicy) bool {
		return policy.RetainIntermediateCommits
	})

	var token *string
	for {
		// Fetch the set of commits for which this upload can resolve code intelligence queries. This
		// will necessarily include the exact commit indicated by teh upload, but may also provide
		// best-effort code intelligence to nearby commits.
		//
		// We need to consider all visible commits, as we may otherwise delete the uploads providing
		// code intelligence for the tip of a branch between the time gitserver is updated and new the
		// associated code intelligence index is processed.
		//
		// We check the set of commits visible to an upload in batches as in some cases it can be very
		// large; for example, a single historic commit providing code intelligence for all descendants.
		commits, nextToken, err := e.dbStore.CommitsVisibleToUpload(ctx, upload.ID, e.commitBatchSize, token)
		if err != nil {
			return false, errors.Wrap(err, "dbstore.CommitsVisibleToUpload")
		}

		if ok, err := e.areCommitsProtectedByPolicy(
			ctx,
			fastPathPolicies,
			slowPathPolicies,
			patterns,
			refDescriptions,
			repositoryCache,
			upload,
			commits,
		); err != nil || ok {
			return ok, err
		}

		if nextToken == nil {
			break
		}

		token = nextToken
	}

	return false, nil
}

func (e *uploadExpirer) areCommitsProtectedByPolicy(ctx context.Context,
	fastPathPolicies []dbstore.ConfigurationPolicy,
	slowPathPolicies []dbstore.ConfigurationPolicy,
	patterns map[string]glob.Glob,
	refDescriptions map[string][]gitserver.RefDescription,
	repositoryCache *repositoryCache,
	upload dbstore.Upload,
	commits []string) (bool, error) {
	for _, commit := range commits {
		// See if this commit was already shown to be protected
		if _, ok := repositoryCache.protectedCommits[commit]; ok {
			return true, nil
		}
	}

	// Try fast path first to avoid another gitserver query. We check _all_ queries in this batch
	// first, as if we find one protected commit we can skip checking all of the others on the slow
	// path as well.
	if e.isUploadProtectedByPolicyFastPath(
		fastPathPolicies,
		patterns,
		refDescriptions,
		repositoryCache,
		upload,
		commits,
	) {
		return true, nil
	}

	// Fall back to slow path
	if ok, err := e.isUploadProtectedByPolicySlowPath(
		ctx,
		slowPathPolicies,
		patterns,
		repositoryCache,
		upload,
		commits,
	); err != nil || ok {
		return ok, err
	}

	return false, nil
}

// isUploadProtectedByPolicyFastPath uses the information we already have about the tips of the repo's
// branches and tags. We will not be able to complete the protection check in this step as we don't yet
// have the data to consider commits contained by a branch, or policies with retain intermediate commits
// enabled. This will be completed in the next step, if the upload is not shown to be protected in this
// "fast path".
func (e *uploadExpirer) isUploadProtectedByPolicyFastPath(
	policies []dbstore.ConfigurationPolicy,
	patterns map[string]glob.Glob,
	refDescriptions map[string][]gitserver.RefDescription,
	repositoryCache *repositoryCache,
	upload dbstore.Upload,
	commits []string,
) bool {
	if len(policies) == 0 {
		return false
	}

	for _, commit := range commits {
		// Match the current working set of policies against the commits, branches, and tags of which
		// the current commit is the tip.
		if ok := newTipPolicyMatcher(patterns, commit, refDescriptions[commit])(policies); ok {
			repositoryCache.protectedCommits[commit] = struct{}{}
			return true
		}
	}

	return false
}

// isUploadProtectedByPolicySlowPath completes the protection check by considering policies with retain
// intermediate commits enabled. Commits contained by a branch are queried from gitserver on demand.
// Gitserver responses are stored in an in-memory LRU cache.
func (e *uploadExpirer) isUploadProtectedByPolicySlowPath(
	ctx context.Context,
	policies []dbstore.ConfigurationPolicy,
	patterns map[string]glob.Glob,
	repositoryCache *repositoryCache,
	upload dbstore.Upload,
	commits []string,
) (bool, error) {
	if len(policies) == 0 {
		return false, nil
	}

	for _, commit := range commits {
		var branches []string
		if v, ok := repositoryCache.branchesContaining.Get(commit); ok {
			branches = v.([]string)
		} else {
			newBranches, err := e.gitserverClient.BranchesContaining(ctx, upload.RepositoryID, commit)
			if err != nil {
				return false, err
			}

			repositoryCache.branchesContaining.Add(commit, newBranches)
			branches = newBranches
		}

		// Match the current working set of policies against the branches of which the current
		// commit belongs. This does not necessarily mean that the branch defines the tip of the
		// branch; that was already checked in the preceding loop.
		if ok := newContainsPolicyMatcher(patterns, commit, branches)(policies); ok {
			repositoryCache.protectedCommits[commit] = struct{}{}
			return true, nil
		}
	}

	return false, nil
}

type repositoryCache struct {
	// protectedCommits is the set of commits that have been shown to be protected. Because we process
	// uploads in descending age, once we write a commit to this map, all future uploads we see visible
	// from this commit will necessarily be younger, and therefore also protected by the same policy.
	protectedCommits map[string]struct{}

	// branchesContaining is an LRU cache from commits to the set of branches that contains that commit.
	// Unfortunately we can't easily order our scan over commits, so it is possible to revisit the same
	// commit at arbitrary intervals, but is unlikely as the order of commits and the order of uploads
	// (which we follow) should usually be correlated. An LRU cache therefore is likely to benefit from
	// some degree of natural locality.
	branchesContaining *lru.Cache
}

func newRepositoryCache(maxKeys int) (*repositoryCache, error) {
	branchesContaining, err := lru.New(maxKeys)
	if err != nil {
		return nil, err
	}

	return &repositoryCache{
		protectedCommits: map[string]struct{}{},
		// TODO - should have "commits unprotected until" map
		branchesContaining: branchesContaining,
	}, nil
}

// filterPolicies returns a new slice containing each of the given policies that pass the given filter.
func filterPolicies(policies []dbstore.ConfigurationPolicy, filter func(policy dbstore.ConfigurationPolicy) bool) []dbstore.ConfigurationPolicy {
	filtered := make([]dbstore.ConfigurationPolicy, 0, len(policies))
	for _, policy := range policies {
		if filter(policy) {
			filtered = append(filtered, policy)
		}
	}

	return filtered
}

// policyCoversUpload returns true if the given policy covers the given upload's age. This function does
// not do any additional checks between teh policy and upload (e.g., target git reference comparisons).
func policyCoversUpload(policy dbstore.ConfigurationPolicy, upload dbstore.Upload) bool {
	return policy.RetentionDuration == nil || timeutil.Now().Sub(*upload.FinishedAt) <= *policy.RetentionDuration
}
