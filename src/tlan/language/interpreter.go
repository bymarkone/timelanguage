package language

import (
	"strconv"
	"tlan/plan"
	"tlan/schedule"
	"tlan/utils"
)

func Eval(context string, items []*Item) {
	switch context {
	case "project":
		evalProject(items)
	case "schedule":
		evalSchedule(items)
	}
}

func evalSchedule(items []*Item) {
	for _, item := range items {
		var track = schedule.Track{}
		track.Schedule = schedule.Schedule{Name: item.Category.Value, Period: findPeriod(item.Category.Annotations, TIME)}
		track.Name = item.Name.Value
		track.Projects = plan.ListProjectsFiltered(func(project plan.Project) bool {
			return project.Category == item.Name.Value
		})
		schedule.AddTrack(track)
	}
}

func evalProject(items []*Item) {
	for _, item := range items {
		var project = plan.Project{}
		project.Name = item.Name.Value
		project.Category = item.Category.Value
		project.Active = !item.Marked
		project.Period = findPeriod(item.Annotations, DATE)
		plan.AddProject(project)
	}
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
	first, _ := strconv.Atoi(binary.Left.Value[0:2])
	second, _ := strconv.Atoi(binary.Left.Value[2:4])
	third, _ := strconv.Atoi(binary.Right.Value[0:2])
	fourth, _ := strconv.Atoi(binary.Right.Value[2:4])
	switch periodType {
	case TIME:
		return utils.Period{Start: utils.DateTime{Hour: first, Minute: second}, End: utils.DateTime{Hour: third, Minute: fourth}}
	case DATE:
		return utils.Period{Start: utils.DateTime{Day: first, Month: second}, End: utils.DateTime{Day: third, Month: fourth}}
	}
	return utils.Period{}
}
