package main

import (
	"fmt"
	"os"
	"os/user"
	"tlan/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to tlan\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}