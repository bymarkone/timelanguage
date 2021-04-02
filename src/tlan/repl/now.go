package repl

import (
	"io"
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
		Function: now,
	}
	RegisterCommands("now", command)
}

func now(_ io.Writer, words []string) {
	tracks := schedule.GetRepository().ListTracks()
	now := time.Now()
	if len(words) > 1 {
		hour, minute := utils.Parse(words[1])
		now = time.Date(now.Year(), now.Month(), now.Day(), hour, minute, now.Second(), now.Nanosecond(), now.Location())
	}
	filteredTracks := schedule.FilterTracks(tracks, func(track schedule.Track) bool {
		return track.Slot.Period.Start.Hour <= now.Hour() && track.Slot.Period.End.Hour > now.Hour() && ContainsWeekday(track.Slot.Period.Weekdays, now.Weekday())
	})
	if len(filteredTracks) == 0 {
		println("It seems you have nothing to do")
		return
	}
	println("NOW is time to do " + filteredTracks[0].Slot.Name)
	for _, track := range filteredTracks {
		subProjects := extractSubProjects(track)
		if len(subProjects) == 0 {
			continue
		}
		println(" ")
		println(track.Name)
		for _, project := range subProjects {
			println(" -- " + project.Name + " [" + project.Parent.Name + "]")
		}
		println(" ")
	}
	printPriorities()
	printDebt()
}

func printPriorities() {
	priorities := planning.GetRepository().GetProject("Priority")
	if priorities == nil {
		return
	}
	if len(priorities.SubProjects) > 0 {
		println("You have also some Priorities")
		for _, project := range priorities.SubProjects {
			println(" -- " + project.Name)
		}
	}
	println(" ")
}

func printDebt() {
	debt := planning.GetRepository().GetProject("Debt")
	if debt == nil {
		return
	}
	if len(debt.SubProjects) > 0 {
		println("And some Debt")
		for _, project := range debt.SubProjects {
			println(" -- " + project.Name)
		}
	}
	println(" ")
}

func extractSubProjects(track *schedule.Track) []*planning.Project {
	var subProjects []*planning.Project
	for _, project := range track.Projects {
		subProjects = append(subProjects, planning.FilterProjects(project.AllSubProjects(), func(item planning.Project) bool {
			return item.Active && (item.Type == "Task" || item.Type == "Pointer")
		})...)
	}
	return subProjects
}

func ContainsWeekday(weekdays []time.Weekday, weekday time.Weekday) bool {
	for _, item := range weekdays {
		if item == weekday {
			return true
		}
	}
	return false
}
