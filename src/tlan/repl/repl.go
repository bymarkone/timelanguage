package repl

import (
	"bufio"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"
	"tlan/language"
	"tlan/planning"
	"tlan/purpose"
	"tlan/schedule"
	"tlan/utils"
)

const Prompt = ">> "
const MaxTableLines = 100

var allCommands = make(map[string]Command)

type Loader struct {
	BaseFolder string
	loaded     map[string]string
}

func (l *Loader) Load() {
	planning.CreateRepository()
	schedule.CreateRepository()
	purpose.CreateRepository()
	l.loaded = make(map[string]string)

	filesInfo, err := ioutil.ReadDir(l.BaseFolder)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, file := range filesInfo {
		//fmt.Printf("Processing file %s \n", file.Name())
		fileAddress := l.BaseFolder + "/" + file.Name()
		content, err := ioutil.ReadFile(fileAddress)
		if err != nil {
			log.Fatal(err)
			return
		}
		context := strings.ReplaceAll(file.Name(), ".gr", "")
		l.loaded[context] = fileAddress
		text := string(content)
		l := language.NewLexer(text)
		p := language.NewParser(l)
		items := p.Parse()
		language.Eval(context, items)
	}
}

func RegisterCommands(name string, command Command) {
	//fmt.Print("Registering '" + name + "' \n")
	allCommands[name] = command
}

var out io.Writer
var loader Loader

func Start(_in io.Reader, _out io.Writer, _loader Loader) {
	scanner := bufio.NewScanner(_in)
	out = _out
	loader = _loader

	for {
		print(Prompt)
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
			allCommands["now"].Function(out, words)
		case "edit":
			allCommands["edit"].Function(out, words)
		case "week":
			allCommands["week"].Function(out, words)
		case "goals":
			allCommands["goals"].Function(out, words)
		}
	}
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
	case "goals":
		printCommand(allCommands[words[1]])
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

	tracks := schedule.GetRepository().ListTracks()
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
	widthMax := 14
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

func slots(_ []string) {
	t := table.NewWriter()
	t.SetOutputMirror(out)

	slots := schedule.GetRepository().ListSlots()

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

	tracks := schedule.GetRepository().ListTracks()

	var header []interface{}
	for _, track := range tracks {
		header = append(header, track.Name+" "+track.Slot.Period.ToString())
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

func printProjects(words []string) {
	inactiveFilter := utils.Find(words, func(val string) bool {
		return val == "--inactive" || val == "-i"
	})
	var projects = planning.GetRepository().ListProjects()
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

func extractFlags(words []string) []string {
	var results []string
	for _, word := range words {
		if strings.HasPrefix(word, "-") {
			results = append(results, strings.ReplaceAll(word, "-", ""))
		}
	}
	return results
}

func hasFlags(flags []string, shallow string) bool {
	return contains(flags, shallow)
}

func contains(flags []string, shallow string) bool {
	for _, item := range flags {
		if item == shallow {
			return true
		}
	}
	return false
}

