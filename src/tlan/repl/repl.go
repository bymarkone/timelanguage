package repl

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
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
		p := language.NewParser(file.Name(), l)
		categories, items := p.Parse()
		language.Eval(context, categories, items)
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
			allCommands["tracks"].Function(out, words)
		case "plan":
			allCommands["plan"].Function(out, words)
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
