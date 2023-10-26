package repl

import (
	"github.com/bymarkone/timelanguage/internal/config"
	"github.com/bymarkone/timelanguage/internal/planning"
	"io"
	"os"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func today(out io.ReadWriter, _ []string) {
	repository := planning.GetRepository()
	tasks := repository.ListTasks()
	types := []string{"Thoughtworks", "Study", "Readings", "Writing", "Coding", "Other"}

	f, err := os.Create(config.EnvBaseFolder() + "/data/today.gr")
	check(err)
	defer func() { check(f.Close()) }()

	_, err = f.WriteString("SELECTED \n\n\n\n")
	check(err)

	for _, taskType := range types {
		printlnint(out, " ")
		_, err := f.WriteString(" \n")
		check(err)
		printlnint(out, taskType)
		_, err = f.WriteString(taskType + " \n")
		check(err)

		selected := planning.FilterTasks(tasks, planning.ByType(taskType))
		selected = planning.FilterTasks(selected, planning.TaskActive)

		for _, task := range selected {
			taskDescription := "- " + task.Project.Name + " :: " + task.Name
			_, err = f.WriteString(taskDescription + " \n")
			check(err)
			printlnint(out, taskDescription)
		}
	}
}
