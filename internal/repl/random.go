package repl

import (
	"github.com/bymarkone/timelanguage/internal/planning"
	"io"
	"math/rand"
)

func init() {
	command := Command{
		Description: "Randomly selects a task from todays list",
		Usage:       "random",
		Arguments:   []Argument{},
		Flags:       []Flag{},
		Function:    random,
	}
	RegisterCommands("random", command)
}

func random(out io.ReadWriter, _ []string) {
	repository := planning.GetRepository()
	tasks := repository.ListTasks()

	tasks = planning.FilterTasks(tasks, planning.TaskActive)

	randomIndex := rand.Intn(len(tasks)) // generate a random int in the range 0 to 9
	pick := tasks[randomIndex]           // get the value from the slice

	taskDescription := "- " + pick.Project.Name + " :: " + pick.Name
	printlnint(out, "The randomly selected task is:")
	printlnint(out, taskDescription)
}
