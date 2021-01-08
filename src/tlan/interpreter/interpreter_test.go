package interpreter

import (
	"testing"
	"tlan/parser"
)

func TestEvalProjects(t *testing.T) {

	tests := []struct {
		input    string
		expected []*Project
	}{
		{
			`
Mathematics
- IU Analysis II
- IU Modern Algebra
- Study Analysis Burkin
- Study Logic for Mathematicians
`,
			[]*Project{
				{Name: "IU Analysis II", Category: "Mathematics"},
				{Name: "IU Modern Algebra", Category: "Mathematics"},
				{Name: "Study Analysis Burkin", Category: "Mathematics"},
				{Name: "Study Logic for Mathematicians", Category: "Mathematics"},
			},
		},
	}

	for _, tt := range tests {

		var context = "project"

		l := parser.NewLexer(tt.input)
		p := parser.NewParser(l)
		items := p.Parse()

		Eval(context, items)
		projects := ListProjects()

		for i, p := range tt.expected {
			if projects[i].Name != p.Name {
				t.Errorf("Project has wrong data2. Got %s, want %s", projects[i].Name, p.Name)
			}
			if projects[i].Category != p.Category {
				t.Errorf("Project has wrong data2. Got %s, want %s", projects[i].Category, p.Category)
			}
		}
	}
}
