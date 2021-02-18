package repl

import (
	"fmt"
	"io"
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
		Function: edit,
	}
	RegisterCommands("edit", command)
}

func edit(out io.Writer ,words []string) {
	if len(words) == 1 {
		cmd := exec.Command("vim", "-O", loader.loaded["goals"], loader.loaded["projects"], loader.loaded["tracks"])
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		loader.Load()
		fmt.Println("Reloaded!")
	} else {
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
}
