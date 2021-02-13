package main

import (
	"fmt"
	"os"
	"os/user"
	"tlan/repl"
)

const DataFolder = "./../../data"

func main() {
	currentUser, err := user.Current()

	if err != nil {
		panic(err)
	}

	loader := repl.Loader{BaseFolder: DataFolder}
	loader.Load()
	fmt.Printf("Hello %s! Welcome to tlan\n", currentUser.Username)
	repl.Start(os.Stdin, os.Stdout, loader)
}
