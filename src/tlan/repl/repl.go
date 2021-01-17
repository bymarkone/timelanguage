package repl

import (
	"bufio"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
	"tlan/planning"
	"tlan/purpose"
	"tlan/schedule"
	"tlan/utils"
)

const PROMPT = ">> "

var out io.Writer

func Start(in io.Reader, _out io.Writer) {
	scanner := bufio.NewScanner(in)
	out = _out

	for {
		print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		words := strings.Split(line, " ")
		if len(words) == 0 {
			continue
		}

		switch words[0] {
		case "show":
			show(words)
		case "help":
			help(words)
		case "clear":
			clear()
		case "exit":
			return
		case "tracks":
			tracks(words)
		case "slots":
			slots(words)
		case "plan":
			plan(words)
		case "now":
			now(words)
		case "edit":
			edit(words)
		case "goals":
			goals(words)
		}
	}
}

func goals(_ []string) {
	t := table.NewWriter()
	t.SetOutputMirror(out)

	goalsByCategory := purpose.GoalsByCategory()

	var header []interface{}
	var headerNames []string
	for category, _ := range goalsByCategory {
		header = append(header, category)
		headerNames = append(headerNames, category)
	}
	t.AppendHeader(header)

	n := 0
	for n < 100 {
		var row []interface{}
		for range headerNames {
			row = append(row, "")
		}
		for i, headerName := range headerNames {
			goals := goalsByCategory[headerName]
			if len(goals) > n && row != nil {
				row[i] = goals[n].Name
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

func edit(words []string) {
	cmd := exec.Command("vim", "./../../data/"+words[1]+"/"+words[2]+".gr")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	fmt.Println(err)
}

func help(words []string) {
	if len(words) == 1 {
		printHelp()
		return
	}

	switch words[1] {
	case "show":
		printShowHelp()
	case "now":
		printNowHelp()
	}
}

func show(words []string) {
	if len(words) == 1 {
		println("Incorrect number of arguments. Type 'help show' to see usage.")
		return
	}
	switch words[1] {
	case "projects":
		printProjects(words)
	}
}

func plan(_ []string) {
	t := table.NewWriter()
	t.SetOutputMirror(out)
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	var header []interface{}
	header = append(header, " ")
	header = append(header, " ")

	currentMonth := time.Now()
	for i := 0; i < 12; i++ {
		header = append(header, currentMonth.Month().String())
		currentMonth = currentMonth.AddDate(0, 1, 0)
	}
	t.AppendHeader(header)

	tracks := schedule.ListTracks()
	for _, track := range tracks {
		now := time.Now()
		var rows []table.Row
		for i := 0; i < 12; i++ {
			var names []string
			for _, project := range track.Projects {
				if project.Period.ActiveIn(now) {
					names = append(names, projectNameForPlan(project))
				}
			}
			for j, name := range names {
				if len(rows) > j {
					if rows != nil {
						rows[j][2+i] = name
					}
				} else {
					var row []interface{}
					row = append(row, track.Slot.Name)
					row = append(row, track.Name)
					for i := 0; i < 12; i++ {
						row = append(row, "")
					}
					row[2+i] = name
					rows = append(rows, row)
				}
			}
			now = now.AddDate(0, 1, 0)
		}
		t.AppendRows(rows, rowConfigAutoMerge)
	}

	t.Style().Options.SeparateRows = true
	widthMax := 15
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: true, WidthMax: widthMax},
		{Number: 2, AutoMerge: true, WidthMax: widthMax},
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
		{Number: 14, WidthMax: widthMax, WidthMaxEnforcer: widthMaxEnforcer},
	})

	t.Render()
}

func widthMaxEnforcer(col string, maxLen int) string {
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
	return name
}

func slots(_ []string) {
	t := table.NewWriter()
	t.SetOutputMirror(out)

	slots := schedule.ListSlots()

	var header []interface{}
	for _, slot := range slots {
		header = append(header, slot.Name)
	}

	n := 0
	for n < 100 {
		var row []interface{}
		for _, slot := range slots {
			row = append(row, extractSlotItemsNames(slot, n))
		}
		if isBlank(row) {
			break
		}
		t.AppendRow(row)
		n++
	}

	t.AppendHeader(header)

	t.Render()
}

func extractSlotItemsNames(slot *schedule.Slot, n int) string {
	if len(slot.FlattenActiveItems(flattener)) >= n+1 {
		return boxedNameForSlots(slot, n)
	}
	return ""
}

func boxedNameForSlots(slot *schedule.Slot, n int) string {
	name := slot.FlattenActiveItems(flattener)[n]
	const LIMIT = 25
	if len(name) > LIMIT {
		return name[0:LIMIT] + "..."
	} else {
		return name
	}
}

func flattener(arr []*schedule.Track) []string {
	return flattenTracksAndProjects(arr)
}

func flattenTracksAndProjects(arr []*schedule.Track) []string {
	var results []string
	for i := range arr {
		results = append(results, strings.ToUpper(arr[i].Name))
		depth := planning.FlattenProjectsDepth(arr[i].Projects)
		results = append(results, toProjectNames(depth)...)
		results = append(results, " ")
	}
	return results
}

func toProjectNames(depth []*planning.Project) []string {
	var results []string
	for i := range depth {
		project := depth[i]
		if project.Active && project.Level == 0 {
			results = append(results, "-"+project.Name)
		}
	}
	return results
}

func tracks(_ []string) {
	print("Tracks:")
	t := table.NewWriter()
	t.SetOutputMirror(out)

	tracks := schedule.ListTracks()

	var header []interface{}
	for _, track := range tracks {
		header = append(header, track.Name)
	}
	t.AppendHeader(header)

	n := 0
	for n < 100 {
		var row []interface{}
		for _, track := range tracks {
			row = append(row, extractProjectNameForTracks(track, n))
		}
		if isBlank(row) {
			break
		}
		t.AppendRow(row)
		n++
	}

	t.Render()
}

func isBlank(row []interface{}) bool {
	for _, item := range row {
		if item != "" {
			return false
		}
	}
	return true
}

func extractProjectNameForTracks(track *schedule.Track, n int) string {
	if len(track.FlattenActiveProjects()) >= n+1 {
		return boxedProjectNameForTracks(track, n)
	}
	return ""
}

func boxedProjectNameForTracks(track *schedule.Track, n int) string {
	project := track.FlattenActiveProjects()[n]
	name := project.Name
	base := ""
	if project.Level >= 1 {
		base = "- " + name
	} else {
		base = name
	}
	const LIMIT = 15
	if len(base) > LIMIT {
		return base[0:LIMIT] + "..."
	} else {
		return base
	}
}

func now(words []string) {
	tracks := schedule.ListTracks()
	now := time.Now()
	if len(words) > 1 {
		hour, minute := utils.Parse(words[1])
		now = time.Date(now.Year(), now.Month(), now.Day(), hour, minute, now.Second(), now.Nanosecond(), now.Location())
	}
	filteredTracks := schedule.FilterTracks(tracks, func(track schedule.Track) bool {
		return track.Slot.Period.Start.Hour <= now.Hour() && track.Slot.Period.End.Hour > now.Hour()
	})
	println("NOW is time to do " + filteredTracks[0].Slot.Name)
	for _, track := range filteredTracks {
		println(track.Name)
		for _, project := range track.Projects {
			if project.Active {
				println(" -- " + project.Name)
			}
		}
	}
}

func printProjects(words []string) {
	inactiveFilter := utils.Find(words, func(val string) bool {
		return val == "--inactive" || val == "-i"
	})
	var projects = planning.ListProjects()
	if inactiveFilter {
		projects = planning.FilterProjects(projects, planning.ByInactive)
	}
	fmt.Print("\nListing projects: \n\n")
	for _, project := range projects {
		fmt.Printf("- %s\n", project.Name)
	}
	fmt.Print("\n\n")
}

func printNowHelp() {
	println("This command prints tasks for a given time slot. Used without arguments, the time that will be considered will be the current time.")
	printEmptyLine()
	println("Usage:")
	println("  now [time]")
	printEmptyLine()
	println("Arguments:")
	println("  time 				                 : time being considered")
}

func printShowHelp() {
	println("This command prints elements of a given collection.")
	printEmptyLine()
	println("Usage:")
	println("  show <collection> {flags}")
	printEmptyLine()
	println("Arguments:")
	println("  collection                  : collection to be printed")
	printEmptyLine()
	println("Flags:")
	println("  --inactive, -i              : display also inactive elements")
}

func printHelp() {
	printEmptyLine()
	printTlanHeader()
	println("Commands:")
	println("  help [command]              : prints help information for commands")
	println("  show <collection>           : prints elements of a given collection")
	println("  now                         : shows tasks to be performed now (i.e the current time slot)")
	println("  exit                        : exits the application")
	printEmptyLine()
}

func println(what string) {
	print(what)
	printEmptyLine()
}

func print(what string) {
	_, _ = fmt.Fprint(out, what)
}

func printEmptyLine() {
	_, _ = fmt.Fprint(out, "\n")
}

func printTlanHeader() {
	println("tLan - a language for time")
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = out
	_ = cmd.Run()
}
