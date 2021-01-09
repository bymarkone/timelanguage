package interpreter

import "tlan/parser"

func Eval(context string, items []*parser.Item) {
	switch context {
	case "project":
		evalProject(items)
	}
}

func evalProject(items []*parser.Item) {
	for _, item := range items {
		var project = Project{}
		project.Name = item.Name.Value
		project.Category = item.Category.Value
		project.Active = !item.Marked
		AddProject(project)
	}
}
