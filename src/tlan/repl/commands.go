package repl

import (
	"text/template"
)

type Command struct {
	Description string
	Usage       string
	Arguments   []Argument
	Flags       []Flag
	function    func([]string)
}

type Argument struct {
	Name        string
	Description string
}

type Flag struct {
	Name        string
	Shortcut    string
	Description string
}

func printCommand(command Command) {
	tmpl, err := template.New("Command").Parse(`
{{.Description}}

Usage:
  {{.Usage}}

Arguments: {{ range $index, $val := .Arguments }}
  {{$val.Name}}                      : {{$val.Description}}{{ end }}
            
Flags: {{ range $index, $val := .Flags }}
  --{{$val.Name}}, -{{$val.Shortcut}}     : {{$val.Description}}{{ end }}
`)

	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(out, command)
	if err != nil {
		panic(err)
	}
}
