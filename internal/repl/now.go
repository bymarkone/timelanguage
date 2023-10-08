package repl

import (
	"github.com/bymarkone/timelanguage/internal/schedule"
	"github.com/bymarkone/timelanguage/internal/utils"
	"io"
	"time"
)

func init() {
	command := Command{
		Description: "Prints what you should be doing now",
		Usage:       "now <hour> <weekday>",
		Arguments: []Argument{
			{Name: "<hour>", Description: "hour to be considered when printing"},
			{Name: "<weekday>", Description: "weekday to be considered when printing"},
		},
		Flags:    []Flag{},
		Function: now,
	}
	RegisterCommands("now", command)
}

func now(out io.ReadWriter, words []string) {
	tracks := schedule.GetRepository().ListTracks()
	now := time.Now()
	if len(words) > 1 {
		hour, minute, _ := utils.Parse(words[1])
		now = time.Date(now.Year(), now.Month(), now.Day(), hour, minute, now.Second(), now.Nanosecond(), now.Location())
	}
	filteredTracks := schedule.FilterTracks(tracks, func(track schedule.Track) bool {
		return track.Slot.Period.Start.Hour <= now.Hour() && track.Slot.Period.End.Hour > now.Hour() && ContainsWeekday(track.Slot.Period.Weekdays, now.Weekday())
	})
	if len(filteredTracks) == 0 {
		printlnint(out, "It seems you have nothing to do")
		return
	}
	printlnint(out, "NOW is time to do "+filteredTracks[0].Slot.Name)
}

func ContainsWeekday(weekdays []time.Weekday, weekday time.Weekday) bool {
	for _, item := range weekdays {
		if item == weekday {
			return true
		}
	}
	return false
}
