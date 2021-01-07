package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if strings.HasPrefix(line, "show") {
			words := strings.Split(line, " ")
			if len(words) == 1 {
				fmt.Print("Incorrect number of arguments\n")
				continue
			}
			switch words[1] {
			case "schedule":
				printSchedule()
			case "plan":
				printPlan()
			}
		} else if strings.ToLower(strings.TrimSpace(line)) == "exit" {
			break
		} else if strings.ToLower(strings.TrimSpace(line)) == "now" {
			fmt.Printf("Hello \n")
		}


	}
}

func printPlan() {
	fmt.Printf("------------------------------------------------------------------------------------------")

}

func printSchedule() {

}

func printPipeline() {

}