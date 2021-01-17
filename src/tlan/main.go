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

const DataFolder = "./../../data"

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
	fmt.Printf("Loading data... \n")
	contexts := []string{"goals", "project", "schedule"}
	for _, context := range contexts {
		loadContext(context)
	}
}
func loadContext(context string) {
	baseFolder := DataFolder + "/" + context
	filesInfo, err := ioutil.ReadDir(baseFolder)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, file := range filesInfo {
		fmt.Printf("Processing file %s \n", file.Name())
		content, err := ioutil.ReadFile(baseFolder + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
			return
		}
		text := string(content)
		l := language.NewLexer(text)
		p := language.NewParser(l)
		items := p.Parse()
		language.Eval(context, items)
	}
}
