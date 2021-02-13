package repl

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	command := Command{
		Description: "This command edit file for objects",
		Usage:       "edit <object>}",
		Arguments: []Argument{
			{Name: "object", Description: "The object to be edited"},
		},
		Flags: []Flag{
			{Name: "", Shortcut: "", Description: ""},
		},
		function: edit,
	}
	registerCommands("edit", command)
}

func edit(words []string) {
	if val, ok := loader.loaded[words[1]]; ok {
		cmd := exec.Command("vim", val)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		loader.Load()
		fmt.Println("Reloaded!")
	}
}
