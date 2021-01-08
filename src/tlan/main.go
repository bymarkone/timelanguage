package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"tlan/parser"
	"tlan/repl"
	"io/ioutil"
	"tlan/interpreter"
)

const DATA_FOLDER = "./../../data"

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	load()
	fmt.Printf("Hello %s! Welcome to tlan\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}

func load() {
	fmt.Printf("Loading data2... \n")
	filesInfo, err := ioutil.ReadDir(DATA_FOLDER)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, file := range filesInfo {
		fmt.Printf("Processing file %s \n", file.Name())
		content, err := ioutil.ReadFile(DATA_FOLDER + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
			return
		}
		text := string(content)
		l := parser.NewLexer(text)
		p := parser.NewParser(l)
		items := p.Parse()
		interpreter.Eval("project", items)
	}
}