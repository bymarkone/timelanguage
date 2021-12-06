package main

import (
	"fmt"
	"github.com/bymarkone/timelanguage/internal/repl"
	"os"
	"os/user"
)

const BaseFolder = "./../.."
const DataFolder = BaseFolder + "/data"
const SamplesFolder = BaseFolder + "/samples"

func main() {
	currentUser, err := user.Current()

	if err != nil {
		panic(err)
	}

	var loader repl.Loader
	if len(os.Args) > 1 && os.Args[1] == "samples" {
		loader = repl.Loader{BaseFolder: SamplesFolder}
	} else {
		loader = repl.Loader{BaseFolder: DataFolder}
	}
	loader.Load()
	fmt.Printf("Hello %s! Welcome to tlan\n", currentUser.Username)
	repl.Start(os.Stdin, os.Stdout, loader)
}
