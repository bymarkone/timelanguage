package repl

import (
	"github.com/bymarkone/timelanguage/internal/schedule"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
)

func init() {
	command := Command{
		Description: "Prints list of tracks and the activities within it",
		Usage:       "tracks",
		Arguments:   []Argument{},
		Flags:       []Flag{},
		Function:    tracks,
	}
	RegisterCommands("tracks", command)
}

func tracks(out io.ReadWriter, _ []string) {
	printlnint(out, "Tracks:")
	t := table.NewWriter()
	t.SetOutputMirror(out)

	tracks := schedule.GetRepository().ListTracks()

	var header []interface{}
	for _, track := range tracks {
		header = append(header, track.Name)
	}
	t.AppendHeader(header)

	var row []interface{}
	for _, track := range tracks {
		row = append(row, track.Slot.Period.ToString())
	}
	t.AppendRow(row)
	t.AppendSeparator()

	n := 0
	for n < 100 {
		var row []interface{}
		for _, track := range tracks {
			row = append(row, extractProjectNameForTracks(track, n))
		}
		if isBlank(row) {
			break
		}
		t.AppendRow(row)
		n++
	}

	t.Render()
}

func isBlank(row []interface{}) bool {
	for _, item := range row {
		if item != "" {
			return false
		}
	}
	return true
}

func extractProjectNameForTracks(track *schedule.Track, n int) string {
	if len(track.FlattenActiveProjects()) >= n+1 {
		return boxedProjectNameForTracks(track, n)
	}
	return ""
}

func boxedProjectNameForTracks(track *schedule.Track, n int) string {
	project := track.FlattenActiveProjects()[n]
	base := project.Name
	const LIMIT = 25
	if len(base) > LIMIT {
		return base[0:LIMIT] + "..."
	} else {
		return base
	}
}
