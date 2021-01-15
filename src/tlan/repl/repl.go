package repl

import (
	"bufio"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
	"os/exec"
	"strings"
	"time"
	"tlan/plan"
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
		case "now":
			now(words)
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
	}
}

func show(words []string) {
	if len(words) == 1 {
		println("Incorrect number of arguments. Type 'help show' to see usage.")
		return
	}
	switch words[1] {
	case "plan":
		printPlan()
	case "projects":
		printProjects(words)
	}
}

func slots(_ []string) {
	t := table.NewWriter()
	t.SetOutputMirror(out)

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

	n := 0
	for n < 100 {
		var row []interface{}
		for _, track := range tracks {
			row = append(row, extractProjectName(track, n))
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

func isBlank(row []interface{}) bool {
	for _, item := range row {
		if item != "" {
			return false
		}
	}
	return true
}

func extractProjectName(track *schedule.Track, n int) string {
	if len(track.FlattenActiveProjects()) >= n+1 {
		return boxedProjectName(track, n)
	}
	return ""
}

func boxedProjectName(track *schedule.Track, n int) string {
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
	var projects = plan.ListProjects()
	if inactiveFilter {
		projects = plan.FilterProjects(projects, plan.ByInactive)
	}
	fmt.Print("\nListing projects: \n\n")
	for _, project := range projects {
		fmt.Printf("- %s\n", project.Name)
	}
	fmt.Print("\n\n")
}

func printPlan() {
	fmt.Printf("------------------------------------------------------------------------------------------")

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
