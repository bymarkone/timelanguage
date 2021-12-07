package language

import (
	"fmt"
	"github.com/bymarkone/timelanguage/internal/planning"
	"github.com/bymarkone/timelanguage/internal/purpose"
	"github.com/bymarkone/timelanguage/internal/schedule"
	"github.com/bymarkone/timelanguage/internal/utils"
	"time"
)

func Eval(context string, categories []*Category, items []*Item) {
	switch context {
	case "projects":
		evalProject(items)
	case "schedule":
		evalSchedule(items)
	case "goals":
		evalGoals(categories, items)
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

func evalGoals(categories []*Category, items []*Item) {
	repository := purpose.GetRepository()
	for _, category := range categories {
		var isChildless = true
		for _, item := range items {
			if category.Value == item.Category.Value {
				isChildless = false
			}
		}
		if isChildless {
			var goal = &purpose.Goal{}
			goal.Name = category.Value
			goal.Category = category.Value
			repository.AddGoal(goal)
		}
	}
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

		repository.AddSlot(slot)

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
	annotation := findUnaryAnnotation(item.Annotations)
	repository := purpose.GetRepository()
	project := &planning.Project{}
	if annotation != nil {
		project.Id = annotation.Name.Value
	}
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
		first, second, third := utils.Parse(binary.Left.Value)
		fourth, fifth, sixth := utils.Parse(binary.Right.Value)
		switch periodType {
		case TIME:
			start = utils.DateTime{Hour: first, Minute: second}
			end = utils.DateTime{Hour: fourth, Minute: fifth}
		case DATE:
			start = utils.DateTime{Day: first, Month: time.Month(second), Year: findYear(third)}
			end = utils.DateTime{Day: fourth, Month: time.Month(fifth), Year: findYear(sixth)}
		}
	}

	return utils.Period{Start: start, End: end, Weekdays: weekdays}
}

func findYear(sixth int) int {
	if sixth == 0 {
		return time.Now().Year()
	} else {
		return sixth
	}
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
