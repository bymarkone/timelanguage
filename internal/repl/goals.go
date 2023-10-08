package repl

import (
	"github.com/bymarkone/timelanguage/internal/planning"
	"github.com/bymarkone/timelanguage/internal/purpose"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"strings"
)

var goalsFlags []string

const GoalsDeep = "deep"

func init() {
	command := Command{
		Description: "Prints the dashboard of goals",
		Usage:       "goals {flags}",
		Arguments:   []Argument{},
		Flags: []Flag{
			{Name: GoalsDeep, Shortcut: "d", Description: "Display projects related to goals"},
		},
		Function: goals,
	}
	RegisterCommands("goals", command)
}

func goals(out io.ReadWriter, words []string) {
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
	isDeep := hasFlags(goalsFlags, GoalsDeep)
	for i := range arr {
		if isDeep {
			results = append(results, strings.ToUpper(arr[i].Name))
			results = append(results, toProjectNamesForGoals(arr[i].Projects)...)
			results = append(results, " ")
		} else {
			results = append(results, arr[i].Name)
		}
	}
	return results
}

func toProjectNamesForGoals(depth []*planning.Project) []string {
	var results []string
	for i := range depth {
		project := depth[i]
		results = append(results, "-"+project.Name)
	
	}
	return results
}
