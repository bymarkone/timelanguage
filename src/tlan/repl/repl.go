package repl

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"tlan/interpreter"
)

const PROMPT = ">> "
var out io.Writer

func Start(in io.Reader, _out io.Writer) {
	scanner := bufio.NewScanner(in)
	out = _out

	for {
		_, _ = fmt.Fprint(out, PROMPT)
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
			break
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
	}
}

func show(words []string) {
	if len(words) == 1 {
		fmt.Print("Incorrect number of arguments. Type 'help show' to see usage.\n")
		return
	}
	switch words[1] {
	case "schedule":
		printSchedule()
	case "plan":
		printPlan()
	case "projects":
		printProjects()
	}
}

func now(words []string) {

}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = out
	_ = cmd.Run()
}

func printProjects() {
	projects := interpreter.ListProjects()
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
	printEmptyLine()
	printTlanHeader()
	_, _ = fmt.Fprint(out, "'show' usage\n")
	printEmptyLine()
}

func printHelp() {
	printEmptyLine()
	printTlanHeader()
	_, _ = fmt.Fprint(out, "Commands:\n")
	_, _ = fmt.Fprint(out, "  help [command]              : prints help information for commands\n")
	_, _ = fmt.Fprint(out, "  show projects|tracks|plan   : prints a specific view for each object\n")
	_, _ = fmt.Fprint(out, "  now                         : shows tasks to be performed now (i.e the current time slot)\n")
	_, _ = fmt.Fprint(out, "  exit                        : exits the application\n")
	printEmptyLine()
}

func printEmptyLine() {
	_, _ = fmt.Fprint(out, "\n")
}

func printTlanHeader() {
	_, _ = fmt.Fprint(out, "tLan - a language for time\n")
}