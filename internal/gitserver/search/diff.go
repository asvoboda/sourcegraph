package search

import (
	"strings"

	git "github.com/libgit2/git2go/v31"
)

// FormattedDiff is a formatted diff between a commit and its parent in the format
// as generated by FormatDiff().
type FormattedDiff string

// FormatDiff generates a formatted diff in the following structure:
//
// oldFile\tnewFile
// @@ hunk header
//  context line
// -removed line
// +added line
func FormatDiff(diff *git.Diff) (FormattedDiff, error) {
	var buf strings.Builder
	buf.Grow(1024)

	err := diff.ForEach(func(delta git.DiffDelta, progress float64) (git.DiffForEachHunkCallback, error) {
		buf.WriteString(delta.OldFile.Path)
		buf.WriteByte('\t')
		buf.WriteString(delta.NewFile.Path)
		buf.WriteByte('\n')

		return func(hunk git.DiffHunk) (git.DiffForEachLineCallback, error) {
			buf.WriteString(hunk.Header)

			return func(line git.DiffLine) error {
				switch line.Origin {
				case git.DiffLineContext:
					buf.WriteByte(' ')
				case git.DiffLineAddition:
					buf.WriteByte('+')
				case git.DiffLineDeletion:
					buf.WriteByte('-')
				default:
					return nil
				}

				buf.WriteString(line.Content)
				return nil
			}, nil
		}, nil
	}, git.DiffDetailLines)

	return FormattedDiff(buf.String()), err
}

// ForEachDelta iterates over the file deltas in a diff in a zero-copy manner
func (d FormattedDiff) ForEachDelta(f func(Delta) bool) {
	remaining := d
	var loc Location
	for len(remaining) > 0 {
		delta := scanDelta(string(remaining))
		remaining = remaining[len(delta):]

		newlineIdx := strings.IndexByte(delta, '\n')
		fileNameLine := delta[:newlineIdx]
		hunks := delta[newlineIdx+1:]
		fileNames := strings.Split(fileNameLine, "\t")
		oldFile, newFile := fileNames[0], fileNames[1]

		if cont := f(Delta{
			location: loc,
			oldFile:  oldFile,
			newFile:  newFile,
			hunks:    hunks,
		}); !cont {
			return
		}

		loc = loc.Shift(Location{
			Offset: len(delta),
			Line:   strings.Count(delta, "\n"),
		})
	}
}

func scanDelta(s string) string {
	offset := 0
	for {
		idx := strings.IndexByte(s[offset:], '\n')
		if idx == -1 {
			return s
		}

		if idx+offset+1 == len(s) {
			return s
		}

		if strings.IndexByte("@+- <>=", s[idx+offset+1]) >= 0 {
			offset += idx + 1
		} else {
			return s[:offset+idx+1]
		}
	}
}

type Delta struct {
	location Location
	oldFile  string
	newFile  string
	hunks    string
}

func (d Delta) OldFile() (string, Location) {
	return d.oldFile, d.location
}

func (d Delta) NewFile() (string, Location) {
	return d.newFile, d.location.Shift(Location{
		Offset: len(d.newFile) + 1,
		Column: len(d.newFile) + 1,
	})
}

// ForEachHunk iterates over each hunk in a delta in a zero-copy manner
func (d Delta) ForEachHunk(f func(Hunk) bool) {
	remaining := d.hunks
	loc := d.location.Shift(Location{Line: 1, Offset: len(d.oldFile) + len(d.newFile) + len(" \n")})
	for len(remaining) > 0 {
		hunk := scanHunk(remaining)
		remaining = remaining[len(hunk):]

		newlineIdx := strings.IndexByte(hunk, '\n')
		header := hunk[:newlineIdx]
		lines := hunk[newlineIdx+1:]

		if cont := f(Hunk{
			location: loc,
			header:   header,
			lines:    lines,
		}); !cont {
			return
		}

		loc = loc.Shift(Location{
			Offset: len(hunk),
			Line:   strings.Count(hunk, "\n"),
		})
	}
}

func scanHunk(s string) string {
	offset := 0
	for {
		idx := strings.IndexByte(s[offset:], '\n')
		if idx == -1 {
			return s
		}

		if idx+offset+1 == len(s) {
			return s
		}

		switch s[idx+offset+1] {
		case '@':
			return s[:offset+idx+1]
		}
		offset += idx + 1
	}
}

type Hunk struct {
	location Location
	header   string
	lines    string
}

func (h Hunk) Header() (string, Location) {
	return h.header, h.location
}

// ForEachLine iterates over each line in a hunk in a zero-copy manner
func (h Hunk) ForEachLine(f func(Line) bool) {
	remaining := h.lines
	loc := h.location.Shift(Location{Line: 1, Offset: len(h.header) + len("\n")})
	for len(remaining) > 0 {
		line := scanLine(remaining)
		remaining = remaining[len(line):]

		if cont := f(Line{
			location: loc,
			fullLine: line,
		}); !cont {
			return
		}

		loc = loc.Shift(Location{
			Offset: len(line),
			Line:   1,
		})
	}
}

func scanLine(s string) string {
	if idx := strings.IndexByte(s, '\n'); idx > 0 {
		return s[:idx+1]
	}
	return s
}

type Line struct {
	location Location
	fullLine string
}

func (l Line) Origin() byte {
	return l.fullLine[0]
}

func (l Line) Content() (string, Location) {
	return l.fullLine[1:], l.location.Shift(Location{Column: 1, Offset: 1})
}
