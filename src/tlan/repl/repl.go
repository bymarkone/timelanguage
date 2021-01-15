package repl

import (
	"bufio"
	"fmt"
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
		case "now":
			now()
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
	}
}

func show(words []string) {
	if len(words) == 1 {
		println("Incorrect number of arguments. Type 'help show' to see usage.")
		return
	}
	switch words[1] {
	case "schedule":
		printSchedule()
	case "plan":
		printPlan()
	case "projects":
		printProjects(words)
	}
}

func now() {
	tracks := schedule.ListTracks()
	now := time.Now()
	filteredTracks := schedule.FilterTracks(tracks, func(track schedule.Track) bool {
		return track.Schedule.Period.Start.Hour < now.Hour() && track.Schedule.Period.End.Hour > now.Hour()
	})
	println("NOW is time to do " + filteredTracks[0].Schedule.Name)
	for _, track := range filteredTracks {
		println(track.Name)
		for _, project := range track.Projects {
			if project.Active {
				println(" -- " + project.Name)
			}
		}
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = out
	_ = cmd.Run()
}

func printProjects(words []string) {
	inactiveFilter := utils.Find(words, func(val string) bool {
		return val == "--inactive" || val == "-i"
	})
	var projects = plan.ListProjects()
	if inactiveFilter {
		projects = plan.FilterProjects(projects, func(val plan.Project) bool {
			return val.Active == false
		})
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

func printSchedule() {

}

func printPipeline() {

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
