package language

import (
	"time"
	"tlan/planning"
	"tlan/purpose"
	"tlan/schedule"
	"tlan/utils"
)

func Eval(context string, items []*Item) {
	switch context {
	case "project":
		evalProject(items)
	case "schedule":
		evalSchedule(items)
	case "goals":
		evalGoals(items)
	}
}

func evalGoals(items []*Item) {
	for _, item := range items {
		var goal = &purpose.Goal{}
		goal.Name = item.Name.Value
		goal.Category = item.Category.Value
		purpose.AddGoal(goal)
	}
}

func evalSchedule(items []*Item) {
	for _, item := range items {
		var track = &schedule.Track{}
		slot := schedule.GetSlot(item.Category.Value)
		if slot == nil {
			slot = &schedule.Slot{Name: item.Category.Value, Period: findPeriod(item.Category.Annotations, TIME)}
			schedule.AddSlot(slot)
		}
		slot.Tracks = append(slot.Tracks, track)
		track.Slot = slot
		track.Name = item.Name.Value
		track.Projects = planning.ListProjectsFiltered(func(project planning.Project) bool {
			return project.Category == item.Name.Value
		})
		schedule.AddTrack(track)
	}
}

func evalProject(items []*Item) {
	for _, item := range items {
		project := projectFromItem(item)
		for _, item := range item.Children {
			subProject := projectFromItem(item)
			subProject.Level = 1
			project.SubProjects = append(project.SubProjects, subProject)
		}
		planning.AddProject(project)
	}
}

func projectFromItem(item *Item) *planning.Project {
	project := &planning.Project{}
	project.Name = item.Name.Value
	project.Category = item.Category.Value
	project.Active = !item.Marked
	project.Period = findPeriod(item.Annotations, DATE)
	if item.Target != "" {
		project.ContributingGoals = append(project.ContributingGoals, item.Target)
		goal := purpose.GetGoal(item.Target)
		if goal == nil {
			goal = purpose.GetGoal(purpose.GoalLess)
		}
		goal.Projects = append(goal.Projects, project)
	}
	return project
}

func findBinaryAnnotation(anns []Annotation) *BinaryAnnotation {
	for _, ann := range anns {
		if ann.Type() == BINARY {
			return ann.(*BinaryAnnotation)
		}
	}
	return nil
}

const (
	TIME = "TIME"
	DATE = "DATE"
)

func findPeriod(anns []Annotation, periodType string) utils.Period {
	binary := findBinaryAnnotation(anns)
	if binary == nil {
		return utils.Period{}
	}
	first, second := utils.Parse(binary.Left.Value)
	third, fourth := utils.Parse(binary.Right.Value)
	switch periodType {
	case TIME:
		return utils.Period{Start: utils.DateTime{Hour: first, Minute: second}, End: utils.DateTime{Hour: third, Minute: fourth}}
	case DATE:
		return utils.Period{Start: utils.DateTime{Day: first, Month: time.Month(second)}, End: utils.DateTime{Day: third, Month: time.Month(fourth)}}
	}
	return utils.Period{}
}
