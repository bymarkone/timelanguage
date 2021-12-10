package repl

import (
	"fmt"
	"github.com/bymarkone/timelanguage/internal/data"
	"golang.org/x/term"
	"io"
	"os"
	"os/exec"
	"strings"
	"unicode/utf8"
)

const MaxTableLines = 100

type TerminalReadWriter struct {
	term *term.Terminal
}

func (t TerminalReadWriter) Write(b []byte) (int, error) {
	return t.term.Write(b)
}

func (t TerminalReadWriter) Read(b []byte) (int, error) {
	return os.Stdin.Read(b)
}

type Repl struct {
	terminal TerminalReadWriter
	data     []string
	loader   data.Loader
}

var allCommands = make(map[string]Command)

func RegisterCommands(name string, command Command) {
	allCommands[name] = command
}

func NewRepl(loader data.Loader) *Repl {
	terminal := term.NewTerminal(os.Stdin, ">")
	return &Repl{
		terminal: TerminalReadWriter{terminal},
		data:     make([]string, 0),
		loader:   loader,
	}
}

func (repl *Repl) ReadInput() (rune, error) {
	var buffer []byte
	for {
		var b = make([]byte, 1)
		_, err := repl.terminal.Read(b)

		if err != nil && err != io.EOF {
			return toRune(b), err
		}

		if b[0] == 3 {
			return toRune(b), nil
		}

		buffer = append(buffer, b[0])

		if len(buffer) > 0 {
			break
		}
	}

	return toRune(buffer), nil
}

func (repl *Repl) Start() {
	var lines []string
	var line string

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	for {
		char, _ := repl.ReadInput()

		switch char {
		case 'q':
			return
		case 13:
			lines = append(lines, line)
			repl.executeCommand(line)
			line = ""
			break
		default:
			str := string(char)
			fmt.Fprint(repl.terminal, str)

			line += str
		}
	}
}

func (repl *Repl) executeCommand(line string) {
	words := strings.Split(line, " ")
	switch words[0] {
	case "clear":
		clear(repl.terminal)
	case "exit":
		return
	case "reload":
		repl.loader.Load()
	case "help":
		allCommands["help"].Function(repl.terminal, words)
	case "tracks":
		allCommands["tracks"].Function(repl.terminal, words)
	case "plan":
		allCommands["plan"].Function(repl.terminal, words)
	case "now":
		allCommands["now"].Function(repl.terminal, words)
	case "edit":
		allCommands["edit"].Function(repl.terminal, words)
	case "week":
		allCommands["week"].Function(repl.terminal, words)
	case "goals":
		allCommands["goals"].Function(repl.terminal, words)
	}
}

func printTlanHeader(out io.ReadWriter) {
	printlnint(out, "tLan - a language for time")
}

func clear(out io.ReadWriter) {
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

func printlnint(out io.ReadWriter, what string) {
	printint(out, what)
	printEmptyLine(out)
}

func printint(out io.ReadWriter, what string) {
	_, _ = fmt.Fprint(out, what)
}

func printEmptyLine(out io.ReadWriter) {
	_, _ = fmt.Fprint(out, "\n")
}

func toRune(b []byte) rune {
	r, _ := utf8.DecodeRune(b)
	return r
}
