package repl

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"strings"
	"time"
	"tlan/planning"
	"tlan/schedule"
)

func init() {
	command := Command{
		Description: "Prints a plan for each month in the next twelve months",
		Usage:       "plan",
		Arguments:   []Argument{},
		Flags:       []Flag{},
		Function:    plan,
	}
	RegisterCommands("plan", command)
}

func plan(out io.Writer, _ []string) {
	t := table.NewWriter()
	t.SetOutputMirror(out)
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	var header []interface{}
	header = append(header, " ")

	currentMonth := time.Now()
	currentMonth = time.Date(currentMonth.Year(), currentMonth.Month(), 1, currentMonth.Hour(),
		currentMonth.Minute(), currentMonth.Second(), currentMonth.Nanosecond(), currentMonth.Location())
	for i := 0; i < 12; i++ {
		header = append(header, currentMonth.Month().String())
		currentMonth = currentMonth.AddDate(0, 1, 0)
	}
	t.AppendHeader(header)

	categoriesWithProjects := planning.GetRepository().ProjectsByCategory()

	now := time.Now()
	now = time.Date(now.Year(), now.Month(), 1, now.Hour(),
		now.Minute(), now.Second(), now.Nanosecond(), now.Location())

	var rows []table.Row
	for _, slot := range schedule.GetRepository().ListSlots() {
		baseRow := len(rows)
		for i := 0; i < 12; i++ {
			var activeProjects []*planning.Project
			for _, track := range schedule.GetRepository().TracksBySlot(*slot) {
				for _, project := range categoriesWithProjects[track.Name] {
					if project.Period.ActiveIn(now) {
						activeProjects = append(activeProjects, project)
					}
				}
			}
			for j, project := range activeProjects {
				if len(rows) > baseRow+j {
					if rows != nil {
						rows[baseRow+j][1+i] = projectNameForPlan(project)
					}
				} else {
					var row []interface{}
					row = append(row, slot.Name)
					for i := 0; i < 12; i++ {
						row = append(row, "")
					}
					row[1+i] = projectNameForPlan(project)
					rows = append(rows, row)
				}
			}
			now = now.AddDate(0, 1, 0)
		}
		now = time.Now()
		now = time.Date(now.Year(), now.Month(), 1, now.Hour(),
			now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	}
	t.AppendRows(rows, rowConfigAutoMerge)

	t.Style().Options.SeparateRows = true
	widthMax := 14
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: true},
		{Number: 2, WidthMax: widthMax, WidthMaxEnforcer: widthMaxEnforcer},
		{Number: 3, WidthMax: widthMax, WidthMaxEnforcer: widthMaxEnforcer},
		{Number: 4, WidthMax: widthMax, WidthMaxEnforcer: widthMaxEnforcer},
		{Number: 5, WidthMax: widthMax, WidthMaxEnforcer: widthMaxEnforcer},
		{Number: 6, WidthMax: widthMax, WidthMaxEnforcer: widthMaxEnforcer},
		{Number: 7, WidthMax: widthMax, WidthMaxEnforcer: widthMaxEnforcer},
		{Number: 8, WidthMax: widthMax, WidthMaxEnforcer: widthMaxEnforcer},
		{Number: 9, WidthMax: widthMax, WidthMaxEnforcer: widthMaxEnforcer},
		{Number: 10, WidthMax: widthMax, WidthMaxEnforcer: widthMaxEnforcer},
		{Number: 11, WidthMax: widthMax, WidthMaxEnforcer: widthMaxEnforcer},
		{Number: 12, WidthMax: widthMax, WidthMaxEnforcer: widthMaxEnforcer},
		{Number: 13, WidthMax: widthMax, WidthMaxEnforcer: widthMaxEnforcer},
	})

	t.Render()
}

func widthMaxEnforcer(col string, _ int) string {
	return strings.Join(strings.Split(col, ","), "\n")
}

func projectNameForPlan(project *planning.Project) string {
	name := project.Name
	const LIMIT = 20
	if len(name) > LIMIT {
		//return name[0:LIMIT] + "."
		return name
	} else {
		return name
	}
}
