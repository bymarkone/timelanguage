package language

import (
	"strconv"
	"tlan/plan"
	"tlan/schedule"
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
		start := findTimeStart(item.Category.Annotations)
		end := findTimeEnd(item.Category.Annotations)
		track.Schedule = schedule.Schedule{Name: item.Category.Value, Period: schedule.Period{Start: start, End: end}}
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
		project.Start = findDayStart(item.Annotations)
		project.End = findDayEnd(item.Annotations)
		plan.AddProject(project)
	}
}

func findDayStart(anns []Annotation) plan.Day {
	for _, ann := range anns {
		if ann.Type() == BINARY {
			binary := ann.(*BinaryAnnotation)
			day, _ := strconv.Atoi(binary.Left.Value[0:2])
			month, _ := strconv.Atoi(binary.Left.Value[2:4])
			return plan.Day{Day: day, Month: month}
		}
	}
	return plan.Day{}
}

func findDayEnd(anns []Annotation) plan.Day {
	for _, ann := range anns {
		if ann.Type() == BINARY {
			binary := ann.(*BinaryAnnotation)
			day, _ := strconv.Atoi(binary.Right.Value[0:2])
			month, _ := strconv.Atoi(binary.Right.Value[2:4])
			return plan.Day{Day: day, Month: month}
		}
	}
	return plan.Day{}
}

func findTimeStart(anns []Annotation) schedule.Time {
	for _, ann := range anns {
		if ann.Type() == BINARY {
			binary := ann.(*BinaryAnnotation)
			day, _ := strconv.Atoi(binary.Left.Value[0:2])
			month, _ := strconv.Atoi(binary.Left.Value[2:4])
			return schedule.Time{Hour: day, Minute: month}
		}
	}
	return schedule.Time{}
}

func findTimeEnd(anns []Annotation) schedule.Time {
	for _, ann := range anns {
		if ann.Type() == BINARY {
			binary := ann.(*BinaryAnnotation)
			day, _ := strconv.Atoi(binary.Right.Value[0:2])
			month, _ := strconv.Atoi(binary.Right.Value[2:4])
			return schedule.Time{Hour: day, Minute: month}
		}
	}
	return schedule.Time{}
}
