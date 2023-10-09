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
	defer f.Close()

	for _, taskType := range types {
		printlnint(out, " ")
		_, err := f.WriteString(" \n")
		check(err)
		printlnint(out, taskType)
		_, err = f.WriteString(taskType + " \n")
		check(err)
		for _, task := range tasks {
			if !(task.Type == taskType) {
				continue
			}
			if task.Active {
				taskDescription := "- " + task.Project.Name + " :: " + task.Name
				_, err = f.WriteString(taskDescription + " \n")
				check(err)
				printlnint(out, taskDescription)
			}
		}
	}

}
