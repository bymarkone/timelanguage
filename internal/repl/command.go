package repl

import (
	"io"
	"text/template"
)

type Command struct {
	Description string
	Usage       string
	Arguments   []Argument
	Flags       []Flag
	Function    func(io.ReadWriter, []string)
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

func printCommand(out io.ReadWriter, command Command) {
	tmpl, err := template.New("Command").Parse(`
{{.Description}}

Usage:
  {{.Usage}}
{{if .Arguments}} Arguments: {{ range $index, $val := .Arguments }}
  {{$val.Name}}                      : {{$val.Description}}{{ end }}
{{end}} 
{{if .Flags}} Flags: {{ range $index, $val := .Flags }}
  --{{$val.Name}}, -{{$val.Shortcut}}     : {{$val.Description}}{{ end }}
{{end}} 
`)

	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(out, command)
	if err != nil {
		panic(err)
	}
}
