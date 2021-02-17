package repl

import (
	"time"
	"tlan/planning"
	"tlan/schedule"
	"tlan/utils"
)

func init() {
	command := Command{
		Description: "This command prints what you should be doing now",
		Usage:       "now <hour> <weekday>",
		Arguments: []Argument{
			{Name: "<hour>", Description: "hour to be considered when printing"},
			{Name: "<weekday>", Description: "weekday to be considered when printing"},
		},
		Flags: []Flag{
			{Name: " ", Shortcut: " ", Description: " "},
		},
		function: now,
	}
	registerCommands("now", command)
}

func now(words []string) {
	tracks := schedule.GetRepository().ListTracks()
	now := time.Now()
	if len(words) > 1 {
		hour, minute := utils.Parse(words[1])
		now = time.Date(now.Year(), now.Month(), now.Day(), hour, minute, now.Second(), now.Nanosecond(), now.Location())
	}
	filteredTracks := schedule.FilterTracks(tracks, func(track schedule.Track) bool {
		return track.Slot.Period.Start.Hour <= now.Hour() && track.Slot.Period.End.Hour > now.Hour() && containsWeekday(track.Slot.Period.Weekdays, now.Weekday())
	})
	if len(filteredTracks) == 0 {
		println("It seems you have nothing to do")
		return
	}
	println("NOW is time to do " + filteredTracks[0].Slot.Name)
	for _, track := range filteredTracks {
		println(track.Name)
		for _, project := range track.Projects {
			for _, subProject := range planning.FilterProjects(project.AllSubProjects(), func(item planning.Project) bool {
				return item.Active && (item.Type == "Task" || item.Type == "Pointer")
			}) {
				println(" -- " + subProject.Name + " [" + project.Name + "]")
			}
		}
		println(" ")
	}
}

func containsWeekday(weekdays []time.Weekday, weekday time.Weekday) bool {
	for _, item := range weekdays {
		if item == weekday {
			return true
		}
	}
	return false
}
