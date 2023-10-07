package main

import (
	"fmt"
	"github.com/bymarkone/timelanguage/internal/config"
	"github.com/bymarkone/timelanguage/internal/data"
	"github.com/bymarkone/timelanguage/internal/repl"
	"os"
	"os/user"
)

func main() {
	var BaseFolder = config.EnvBaseFolder()
	var DataFolder = BaseFolder + "/data"
	var SamplesFolder = BaseFolder + "/samples"

	currentUser, err := user.Current()

	if err != nil {
		panic(err)
	}

	var loader data.Loader
	if len(os.Args) > 1 && os.Args[1] == "samples" {
		loader = data.Loader{BaseFolder: SamplesFolder}
	} else {
		loader = data.Loader{BaseFolder: DataFolder}
	}

	loader.Load()
	fmt.Printf("Hello %s! Welcome to tlan\n", currentUser.Username)
	//repl.Start(os.Stdin, os.Stdout, loader)

	_repl := repl.NewRepl(loader)

	_repl.Start()
}

func loadConfiguration() {
}
