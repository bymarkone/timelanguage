package repl

import (
	"github.com/jedib0t/go-pretty/table"
	"strings"
	"tlan/planning"
	"tlan/purpose"
)

var goalsFlags []string

const GoalsShallow = "shallow"

func goals(words []string) {
	t := table.NewWriter()
	t.SetOutputMirror(out)

	goalsFlags = extractFlags(words)

	goalsByCategory := purpose.GoalsByCategory()

	var header []interface{}
	var headerNames []string
	for category := range goalsByCategory {
		header = append(header, category)
		headerNames = append(headerNames, category)
	}
	t.AppendHeader(header)

	n := 0
	for n < MaxTableLines {
		row := make([]interface{}, len(headerNames))
		for i, headerName := range headerNames {
			names := flattenGoalsAndProjects(goalsByCategory[headerName])
			if len(names) > n {
				row[i] = names[n]
			} else {
				row[i] = ""
			}
		}
		if isBlank(row) {
			break
		}
		t.AppendRow(row)
		n++
	}

	t.Render()
}

func flattenGoalsAndProjects(arr []*purpose.Goal) []string {
	var results []string
	isShallow := hasFlags(goalsFlags, GoalsShallow)
	for i := range arr {
		if isShallow {
			results = append(results, arr[i].Name)
		} else {
			results = append(results, strings.ToUpper(arr[i].Name))
			results = append(results, toProjectNamesForGoals(arr[i].Projects)...)
		}
		results = append(results, " ")
	}
	return results
}

func toProjectNamesForGoals(depth []*planning.Project) []string {
	var results []string
	for i := range depth {
		project := depth[i]
		if project.Level == 0 {
			results = append(results, "-"+project.Name)
		}
	}
	return results
}
