package repl

import (
	"io"
)

func init() {
	command := Command{
		Description: "Edit file for objects",
		Usage:       "edit <object>",
		Arguments: []Argument{
			{Name: "object", Description: "The object to be edited"},
		},
		Flags: []Flag{ },
		Function: edit,
	}
	RegisterCommands("edit", command)
}

func edit(_ io.ReadWriter ,_ []string) {
}
