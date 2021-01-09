package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"tlan/language"
	"tlan/repl"
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
		l := language.NewLexer(text)
		p := language.NewParser(l)
		items := p.Parse()
		language.Eval("project", items)
	}
}