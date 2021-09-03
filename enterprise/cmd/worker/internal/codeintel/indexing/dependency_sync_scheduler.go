package indexing

import (
	"context"
	"strconv"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/hashicorp/go-multierror"
	"github.com/inconshreveable/log15"

	"github.com/sourcegraph/sourcegraph/enterprise/internal/codeintel/stores/dbstore"
	"github.com/sourcegraph/sourcegraph/internal/actor"
	"github.com/sourcegraph/sourcegraph/internal/database"
	"github.com/sourcegraph/sourcegraph/internal/extsvc"
	"github.com/sourcegraph/sourcegraph/internal/observation"
	"github.com/sourcegraph/sourcegraph/internal/workerutil"
	"github.com/sourcegraph/sourcegraph/internal/workerutil/dbworker"
	dbworkerstore "github.com/sourcegraph/sourcegraph/internal/workerutil/dbworker/store"
	"github.com/sourcegraph/sourcegraph/lib/codeintel/precise"
)

var schemeToExternalService = map[string]string{
	"semanticdb": extsvc.KindJVMPackages,
}

// NewDependencySyncScheduler returns a new worker instance that processes
// records from lsif_dependency_indexing_jobs.
func NewDependencySyncScheduler(
	dbStore DBStore,
	workerStore dbworkerstore.Store,
	externalServiceStore ExternalServiceStore,
	metrics workerutil.WorkerMetrics,
) *workerutil.Worker {
	rootContext := actor.WithActor(context.Background(), &actor.Actor{Internal: true})

	handler := &dependencySyncSchedulerHandler{
		dbStore:     dbStore,
		workerStore: workerStore,
		extsvcStore: externalServiceStore,
	}

	return dbworker.NewWorker(rootContext, workerStore, handler, workerutil.WorkerOptions{
		Name:              "precise_code_intel_dependency_sync_scheduler_worker",
		NumHandlers:       1,
		Interval:          time.Second * 5,
		HeartbeatInterval: 1 * time.Second,
		Metrics:           metrics,
	})
}

type dependencySyncSchedulerHandler struct {
	dbStore     DBStore
	workerStore dbworkerstore.Store
	extsvcStore ExternalServiceStore
}

func (h *dependencySyncSchedulerHandler) Handle(ctx context.Context, record workerutil.Record) error {
	job := record.(dbstore.DependencyIndexingJob)

	log15.Info("GOT NEW INDEXING JOB")

	scanner, err := h.dbStore.ReferencesForUpload(ctx, job.UploadID)
	if err != nil {
		return errors.Wrap(err, "dbstore.ReferencesForUpload")
	}
	defer func() {
		if closeErr := scanner.Close(); closeErr != nil {
			err = multierror.Append(err, errors.Wrap(closeErr, "dbstore.ReferencesForUpload.Close"))
		}
	}()

	var (
		kinds                      = map[string]struct{}{}
		oldDependencyReposInserted int
		newDependencyReposInserted int
	)
	var errs []error

	for {
		packageReference, exists, err := scanner.Next()
		if err != nil {
			return errors.Wrap(err, "dbstore.ReferencesForUpload.Next")
		}
		if !exists {
			break
		}

		pkg := precise.Package{
			Scheme:  packageReference.Package.Scheme,
			Name:    packageReference.Package.Name,
			Version: packageReference.Package.Version,
		}

		extsvcKind, ok := schemeToExternalService[packageReference.Scheme]
		kinds[extsvcKind] = struct{}{}
		if !ok {
			continue
		}

		new, err := h.insertDependencyRepo(ctx, pkg)
		if err != nil {
			errs = append(errs, err)
		} else if new {
			newDependencyReposInserted++
		} else {
			oldDependencyReposInserted++
		}
	}

	var nextSync *time.Time
	// If len == 0, it will return all external services, which we definitely don't want.
	if len(kinds) > 0 {
		nextSync = timePtr(time.Now())
		externalServices, err := h.extsvcStore.List(ctx, database.ExternalServicesListOptions{
			Kinds: kindsToArray(kinds),
		})
		if err != nil {
			if len(errs) == 0 {
				return errors.Wrap(err, "dbstore.List")
			} else {
				return multierror.Append(err, errs...)
			}
		}

		log15.Info("syncing external services",
			"upload", job.UploadID, "num", len(externalServices), "job", job.ID, "schemaKinds", kinds,
			"newRepos", newDependencyReposInserted, "existingInserts", oldDependencyReposInserted)

		for _, externalService := range externalServices {
			externalService.NextSyncAt = *nextSync
			err := h.extsvcStore.Upsert(ctx, externalService)
			if err != nil {
				errs = append(errs, errors.Wrapf(err, "extsvcStore.Upsert: error setting next_sync_at for external service %d - %s", externalService.ID, externalService.DisplayName))
			}
		}

	} else {
		log15.Info("no package schema kinds to sync external services for", "upload", job.UploadID, "job", job.ID)
	}

	// append empty kind as queueing jobs are partitioned on extsvc kind, and we want queueing jobs for
	// uploads not associated with explicitly syncing an external service e.g. Go uploads
	for kind := range kinds {
		if _, err := h.dbStore.InsertDependencyIndexingQueueingJob(ctx, job.UploadID, kind, nextSync); err != nil {
			errs = append(errs, errors.Wrap(err, "dbstore.InsertDependencyIndexingQueueingJob"))
		}
	}

	if len(errs) == 0 {
		return nil
	}

	if len(errs) == 1 {
		return errs[0]
	}

	return multierror.Append(nil, errs...)
}

func (h *dependencySyncSchedulerHandler) insertDependencyRepo(ctx context.Context, pkg precise.Package) (new bool, err error) {
	ctx, endObservation := dependencyReposOps.InsertCloneableDependencyRepo.With(ctx, &err, observation.Args{
		MetricLabelValues: []string{pkg.Scheme},
	})
	defer func() {
		endObservation(1, observation.Args{MetricLabelValues: []string{strconv.FormatBool(new)}})
	}()

	new, err = h.dbStore.InsertCloneableDependencyRepo(ctx, pkg)
	if err != nil {
		return new, errors.Wrap(err, "dbstore.InsertCloneableDependencyRepos")
	}
	return new, nil
}

func timePtr(t time.Time) *time.Time { return &t }

func kindsToArray(k map[string]struct{}) (s []string) {
	for kind := range k {
		if kind != "" {
			s = append(s, kind)
		}
	}
	return
}
