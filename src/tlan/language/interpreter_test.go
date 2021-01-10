package language

import (
	"testing"
	"tlan/plan"
)

func TestEvalProjects(t *testing.T) {

	tests := []struct {
		input    string
		expected []*plan.Project
	}{
		{
			`
Mathematics
- IU Analysis II
- IU Modern Algebra [1001-1504]
- Study Analysis Burkin
- (Study Logic for Mathematicians)
`,
			[]*plan.Project{
				{Name: "IU Analysis II", Category: "Mathematics", Active: true},
				{Name: "IU Modern Algebra", Category: "Mathematics", Active: true, Start: plan.Day{Day: 10, Month: 1}, End: plan.Day{Day: 15, Month: 4}},
				{Name: "Study Analysis Burkin", Category: "Mathematics", Active: true},
				{Name: "Study Logic for Mathematicians", Category: "Mathematics", Active: false},
			},
		},
	}

	for _, tt := range tests {

		var context = "project"

		l := NewLexer(tt.input)
		p := NewParser(l)
		items := p.Parse()

		Eval(context, items)
		projects := plan.ListProjects()

		for i, p := range tt.expected {
			if projects[i].Name != p.Name {
				t.Errorf("Project has wrong data. Got %s, want %s", projects[i].Name, p.Name)
			}
			if projects[i].Category != p.Category {
				t.Errorf("Project has wrong data. Got %s, want %s", projects[i].Category, p.Category)
			}
			if projects[i].Active != p.Active {
				t.Errorf("Project has wrong attribute. Got %v, want %v", projects[i].Active, p.Active)
			}
			if projects[i].Start != p.Start {
				t.Errorf("Project has wrong attribute. Got %v, want %v", projects[i].Start, p.Start)
			}
		}
	}
}
