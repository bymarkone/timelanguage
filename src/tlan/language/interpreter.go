package language

import (
	"fmt"
	"time"
	"tlan/planning"
	"tlan/purpose"
	"tlan/schedule"
	"tlan/utils"
)

func Eval(context string, items []*Item) {
	switch context {
	case "projects":
		evalProject(items)
	case "schedule":
		evalSchedule(items)
	case "goals":
		evalGoals(items)
	case "tasks":
		evalTasks(items)
	}
}

func evalTasks(items []*Item) {
	repository := planning.GetRepository()
	for _, item := range items {
		var project = projectFromItem(item)
		parent := repository.GetProject(item.Category.Value)
		if parent == nil {
			parent = &planning.Project{}
			parent.Name = item.Category.Value
			repository.AddProject(parent)
		}
		project.Parent = parent
		parent.SubProjects = append(parent.SubProjects, project)
	}
}

func evalGoals(items []*Item) {
	repository := purpose.GetRepository()
	for _, item := range items {
		var goal = &purpose.Goal{}
		goal.Name = item.Name.Value
		goal.Category = item.Category.Value
		repository.AddGoal(goal)
	}
}

func evalSchedule(items []*Item) {
	repository := schedule.GetRepository()
	for _, item := range items {
		var track = &schedule.Track{}

		slotPeriod := findPeriod(item.Category.Annotations, TIME, utils.Period{Weekdays: allWeekDays})
		trackPeriod := findPeriod(item.Annotations, TIME, slotPeriod)
		slot := &schedule.Slot{Name: item.Category.Value, Period: trackPeriod}

		slot.Tracks = append(slot.Tracks, track)
		track.Slot = slot
		track.Name = item.Name.Value
		track.Projects = planning.GetRepository().ListProjectsFiltered(func(project planning.Project) bool {
			return project.Category == item.Name.Value
		})
		repository.AddTrack(track)
	}
}

func evalProject(items []*Item) {
	repository := planning.GetRepository()
	for _, item := range items {
		project := projectFromItem(item)
		for _, item := range item.Children {
			subProject := projectFromItem(item)
			subProject.Level = 1
			subProject.Parent = project
			project.SubProjects = append(project.SubProjects, subProject)
		}
		repository.AddProject(project)
	}
}

func projectFromItem(item *Item) *planning.Project {
	repository := purpose.GetRepository()
	project := &planning.Project{}
	project.Name = item.Name.Value
	project.Category = item.Category.Value
	project.Active = !item.Marked
	project.Period = findPeriod(item.Annotations, DATE, utils.Period{})
	project.Type = item.Type
	if item.Target != "" {
		project.ContributingGoals = append(project.ContributingGoals, item.Target)
		goal := repository.GetGoal(item.Target)
		if goal == nil {
			fmt.Printf("Invalid goal for project %s \n", project.Name)
		} else {
			goal.Projects = append(goal.Projects, project)
		}
	}
	return project
}

const (
	TIME = "TIME"
	DATE = "DATE"
)

func findPeriod(anns []Annotation, periodType string, parentPeriod utils.Period) utils.Period {
	binary := findBinaryAnnotation(anns)
	unary := findUnaryAnnotation(anns)

	weekdays := parentPeriod.Weekdays
	start := parentPeriod.Start
	end := parentPeriod.End

	if unary != nil {
		weekdays = computeWeekdays(unary.Name.Value)
	}

	if binary != nil {
		first, second := utils.Parse(binary.Left.Value)
		third, fourth := utils.Parse(binary.Right.Value)
		switch periodType {
		case TIME:
			start = utils.DateTime{Hour: first, Minute: second}
			end = utils.DateTime{Hour: third, Minute: fourth}
		case DATE:
			start = utils.DateTime{Day: first, Month: time.Month(second)}
			end = utils.DateTime{Day: third, Month: time.Month(fourth)}
		}
	}

	return utils.Period{Start: start, End: end, Weekdays: weekdays}
}

var shortWeekdays = map[string]time.Weekday{
	"Mon": time.Monday,
	"Tue": time.Tuesday,
	"Wed": time.Wednesday,
	"Thu": time.Thursday,
	"Fri": time.Friday,
	"Sat": time.Saturday,
	"Sun": time.Sunday,
}

var allWeekDays = []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}

func computeWeekdays(pattern string) []time.Weekday {
	var weekdays []time.Weekday
	switch pattern {
	case "Daily":
		weekdays = allWeekDays
	default:
		res := ""
		for i, r := range pattern {
			res = res + string(r)
			if i > 0 && (i+1)%3 == 0 {
				weekdays = append(weekdays, shortWeekdays[res])
				res = ""
			}
		}
	}
	return weekdays
}

func findUnaryAnnotation(anns []Annotation) *UnaryAnnotation {
	for _, ann := range anns {
		if ann.Type() == UNARY {
			return ann.(*UnaryAnnotation)
		}
	}
	return nil
}

func findBinaryAnnotation(anns []Annotation) *BinaryAnnotation {
	for _, ann := range anns {
		if ann.Type() == BINARY {
			return ann.(*BinaryAnnotation)
		}
	}
	return nil
}
