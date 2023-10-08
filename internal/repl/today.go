package repl

import (
	"github.com/bymarkone/timelanguage/internal/planning"
	"io"
)

func init() {
	command := Command{
		Description: "Shows which tasks are available for today",
		Usage:       "today",
		Arguments:   []Argument{},
		Flags:       []Flag{},
		Function:    today,
	}
	RegisterCommands("today", command)
}

func today(out io.ReadWriter, _ []string) {
	repository := planning.GetRepository()
	tasks := repository.ListTasks()

	for _, task := range tasks {
		if task.Active {
			printlnint(out, "- "+task.Project.Name+" :: "+task.Name)
		}
	}

}
