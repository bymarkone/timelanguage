package language

import (
	"testing"
	"tlan/plan"
	"tlan/schedule"
	"tlan/utils"
)

func TestEvalTracks(t *testing.T) {

	period := utils.Period{Start: utils.DateTime{Hour: 5, Minute: 0}, End: utils.DateTime{Hour: 9, Minute: 0}}
	expectedSchedule := schedule.Schedule{Name: "Creative Work", Period: period}

	tests := []struct {
		input    string
		expected []*schedule.Track
	}{
		{
			`
Creative Work [0500-0900]
* Mathematics
* Books
* Research
* Bible
`,
			[]*schedule.Track{
				{Name: "Mathematics", Schedule: expectedSchedule},
				{Name: "Books", Schedule: expectedSchedule},
				{Name: "Research", Schedule: expectedSchedule},
				{Name: "Bible", Schedule: expectedSchedule},
			},
		},
	}

	schedule.Clean()

	for _, tt := range tests {

		l := NewLexer(tt.input)
		p := NewParser(l)
		items := p.Parse()

		Eval("schedule", items)
		tracks := schedule.ListTracks()

		for i, r := range tt.expected {
			if tracks[i].Name != r.Name {
				t.Errorf("Track has wrong data. Got %s, want %s", tracks[i].Name, r.Name)
			}
			if tracks[i].Schedule != r.Schedule {
				t.Errorf("Track has wrong data. Got %v, want %v", tracks[i].Schedule, r.Schedule)
			}
		}
	}
}

func TestEvalProjects(t *testing.T) {

	period := utils.Period{Start: utils.DateTime{Day: 10, Month: 1}, End: utils.DateTime{Day: 15, Month: 4}}

	tests := []struct {
		input    string
		expected []*plan.Project
	}{
		{
			`
Mathematics
- IU Analysis II >> BS Mathematics
- IU Modern Algebra [1001-1504]
* Study Analysis Burkin
- (Study Logic for Mathematicians)
`,
			[]*plan.Project{
				{Name: "IU Analysis II", Category: "Mathematics", Active: true, ContributingGoals: []*plan.Goal{{"BS Mathematics"}}},
				{Name: "IU Modern Algebra", Category: "Mathematics", Active: true, Period: period, ContributingGoals: []*plan.Goal{}},
				{Name: "Study Analysis Burkin", Category: "Mathematics", Active: true, ContributingGoals: []*plan.Goal{}},
				{Name: "Study Logic for Mathematicians", Category: "Mathematics", Active: false, ContributingGoals: []*plan.Goal{}},
			},
		},
	}

	plan.Clean()

	for _, tt := range tests {

		l := NewLexer(tt.input)
		p := NewParser(l)
		items := p.Parse()

		Eval("project", items)
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
			if projects[i].Period.Start != p.Period.Start {
				t.Errorf("Project has wrong attribute. Got %v, want %v", projects[i].Period.Start, p.Period.Start)
			}
			if !equal(projects[i].ContributingGoals, p.ContributingGoals) {
				t.Errorf("Project has wrong goals. Got %v, want %v", projects[i].ContributingGoals, p.ContributingGoals)
			}
		}
	}
}

func equal(first []*plan.Goal, second []*plan.Goal) bool {
	if len(first) != len(second) {
		return false
	}
	for i, _ := range first {
		if !(*first[i] == *second[i]) {
			return false
		}
	}
	return true
}
