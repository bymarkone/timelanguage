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
const Prompt = ">> "

type TerminalReadWriter struct {
	term *term.Terminal
}

func (t TerminalReadWriter) Write(b []byte) (int, error) {
	return t.term.Write(b)
}

func (t TerminalReadWriter) Read(b []byte) (int, error) {
	return os.Stdin.Read(b)
}

func (t TerminalReadWriter) Print(text string) {
	_, err := fmt.Fprint(t.term, text)
	if err != nil {
		print("Error printing to terminal")
	}
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

func (repl *Repl) ReadInput() rune {
	var buffer []byte
	for {
		var b = make([]byte, 1)
		_, err := repl.terminal.Read(b)

		if err != nil && err != io.EOF {
			return toRune(b)
		}

		if b[0] == 3 {
			return toRune(b)
		}

		buffer = append(buffer, b[0])

		if len(buffer) > 0 {
			break
		}
	}

	return toRune(buffer)
}

func (repl *Repl) Start() {
	var lines []string
	var line string
	var col = 0
	var row = 0

	terminal := repl.terminal


	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	terminal.Print(Prompt)

	for {
		char := repl.ReadInput()

		switch char {
		case 'q':
			return
		case 127:
			if col == 0 {
				break
			}
			col -= 1
			terminal.Print("\b \b")
			line = line[:len(line)-1]
			break
		case 13:
			if len(line) > 0 {
				lines = append(lines, line)
				repl.executeCommand(line)
				line = ""
			}
			terminal.Print("\n")
			terminal.Print(Prompt)
			col = 0
			break
		case 27:
			repl.ReadInput()
			arrow := repl.ReadInput()
			switch arrow {
			case 68:
				if col == 1 {
					break
				}
				terminal.Print("\b")
				col -= 1
				break
			case 65:
				index := len(lines) - (row+1)
				if index < 0 || index > len(lines)-1 {
					break
				}
				row += 1
				terminal.Print("\r")
				terminal.Print(strings.Repeat(" ", len(line) + 2))
				terminal.Print("\r")
				line = lines[index]
				terminal.Print(Prompt)
				terminal.Print(line)
				break
			case 66:
				index := len(lines) - (row-1)
				 if index < 0 || index > len(lines)-1 {
					break
				}
				row -= 1
				terminal.Print("\r")
				terminal.Print(strings.Repeat(" ", len(line) + 2))
				terminal.Print("\r")
				line = lines[index]
				terminal.Print(Prompt)
				terminal.Print(line)
				break
			case 67:
				if col == len(line) {
					break
				}
				terminal.Print(string(line[col]))
				col += 1
				break
			}
			break
		default:
			str := string(char)
			terminal.Print(str)

			line += str
			col += 1
		}
	}
}

func (repl *Repl) executeCommand(line string) {
	words := strings.Split(line, " ")
	terminal := repl.terminal

	switch words[0] {
	case "clear":
		clear(repl.terminal)
	case "exit":
		return
	case "reload":
		repl.loader.Load()
	case "help":
		terminal.Print("\n")
		allCommands["help"].Function(repl.terminal, words)
	case "tracks":
		terminal.Print("\n")
		allCommands["tracks"].Function(repl.terminal, words)
	case "plan":
		terminal.Print("\n")
		allCommands["plan"].Function(repl.terminal, words)
	case "now":
		terminal.Print("\n")
		allCommands["now"].Function(repl.terminal, words)
	case "edit":
		terminal.Print("\n")
		allCommands["edit"].Function(repl.terminal, words)
	case "week":
		terminal.Print("\n")
		allCommands["week"].Function(repl.terminal, words)
	case "goals":
		terminal.Print("\n")
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
