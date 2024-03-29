package repl

import (
	"fmt"
	"io"
)

func init() {
	command := Command{
		Description: "Prints help for tLan commands",
		Usage:       "help {command}",
		Arguments: []Argument{
			{Name: "command", Description: "[Optional] The command for which help will be displayed"},
		},
		Flags:    []Flag{},
		Function: help,
	}
	RegisterCommands("help", command)
}

func help(out io.ReadWriter, words []string) {
	if len(words) == 1 {
		printHelp(out)
		return
	} else {
		command := allCommands[words[1]]
		printCommand(out, command)
	}

}

func printHelp(out io.ReadWriter) {
	printEmptyLine(out)
	printTlanHeader(out)
	printlnint(out, "Commands:")
	format := "  %-27v : %s\n"
	for _, command := range allCommands {
		fmt.Fprintf(out, format, command.Usage, command.Description)
	}
	fmt.Fprintf(out, format, "exit", "Exits the application")
	printEmptyLine(out)
}
