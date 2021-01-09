package language

import (
	"tlan/plan"
)

func Eval(context string, items []*Item) {
	switch context {
	case "project":
		evalProject(items)
	}
}

func evalProject(items []*Item) {
	for _, item := range items {
		var project = plan.Project{}
		project.Name = item.Name.Value
		project.Category = item.Category.Value
		project.Active = !item.Marked
		plan.AddProject(project)
	}
}
