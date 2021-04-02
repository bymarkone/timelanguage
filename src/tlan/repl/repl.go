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
		case "clear":
			clear()
		case "exit":
			return
		case "help":
			allCommands["help"].Function(out, words)
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
