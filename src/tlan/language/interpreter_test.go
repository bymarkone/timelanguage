package language

import (
	"reflect"
	"testing"
	"time"
	"tlan/planning"
	"tlan/purpose"
	"tlan/schedule"
	"tlan/utils"
)

func TestEvalGoals(t *testing.T) {
	tests := []struct {
		input    string
		expected []*purpose.Goal
	}{
		{
			`
Great Technologist [Lagging]
- Drive strategy and execution [Leading]
- First class engineer [Leading]
- Democratize best practices [Leading]
`,
			[]*purpose.Goal{
				{Name: "Drive strategy and execution", Tags: []string{"Leading"}, Category: "Great Technologist"},
				{Name: "First class engineer", Tags: []string{"Leading"}, Category: "Great Technologist"},
				{Name: "Democratize best practices", Tags: []string{"Leading"}, Category: "Great Technologist"},
			},
		},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		p := NewParser(l)
		items := p.Parse()

		Eval("goals", items)
		goals := purpose.GetRepository().ListGoals()

		for i, g := range tt.expected {
			if goals[i].Name != g.Name {
				t.Errorf("Goal has the wrong name. Got %s, want %s", goals[i].Name, g.Name)
			}
			if goals[i].Category != g.Category {
				t.Errorf("Goal has the wrong category. Got %s, want %s", goals[i].Category, g.Category)
			}
		}
	}

}

func TestEvalTracks(t *testing.T) {

	period := utils.Period{Start: utils.DateTime{Hour: 5, Minute: 0}, End: utils.DateTime{Hour: 9, Minute: 0}, Weekdays: allWeekDays}
	expectedSchedule := &schedule.Slot{Name: "Creative Work", Period: period}
	periodForBooks := utils.Period{Start: utils.DateTime{Hour: 5, Minute: 0}, End: utils.DateTime{Hour: 9, Minute: 0}, Weekdays: []time.Weekday{time.Monday, time.Tuesday}}
	expectedScheduleForBook := &schedule.Slot{Name: "Creative Work", Period: periodForBooks}

	tests := []struct {
		input    string
		expected []*schedule.Track
	}{
		{
			`
Creative Work [Daily, 05:00-09:00]
* Mathematics
* Books [MonTue]
* Research
* Bible
`,
			[]*schedule.Track{
				{Name: "Mathematics", Slot: expectedSchedule},
				{Name: "Books", Slot: expectedScheduleForBook},
				{Name: "Research", Slot: expectedSchedule},
				{Name: "Bible", Slot: expectedSchedule},
			},
		},
	}

	schedule.CreateRepository()

	for _, tt := range tests {

		l := NewLexer(tt.input)
		p := NewParser(l)
		items := p.Parse()

		Eval("schedule", items)
		tracks := schedule.GetRepository().ListTracks()

		for i, r := range tt.expected {
			if tracks[i].Name != r.Name {
				t.Errorf("Track has wrong data. Got %s, want %s", tracks[i].Name, r.Name)
			}
			if tracks[i].Slot.Period.Start != r.Slot.Period.Start {
				t.Errorf("Track has wrong slot data. Got %v, want %v", tracks[i].Slot, r.Slot)
			}
			if tracks[i].Slot.Period.End != r.Slot.Period.End {
				t.Errorf("Track has wrong slot data. Got %v, want %v", tracks[i].Slot, r.Slot)
			}
			if !reflect.DeepEqual(tracks[i].Slot.Period.Weekdays, r.Slot.Period.Weekdays) {
				t.Errorf("Track has wrong slot data. Got %v, want %v", tracks[i].Slot, r.Slot)
			}
		}
	}
}

func TestEvalProjects(t *testing.T) {

	period := utils.Period{Start: utils.DateTime{Day: 10, Month: 1}, End: utils.DateTime{Day: 15, Month: 4}}

	tests := []struct {
		input    string
		expected []*planning.Project
	}{
		{
			`
Mathematics
- IU Analysis II >> BS Mathematics
- IU Modern Algebra [10/01-15/04]
  * Read book
- Study Analysis Burkin
  + Follow another list
- (Study Logic for Mathematicians)
`,
			[]*planning.Project{
				{Name: "IU Analysis II", Category: "Mathematics", Active: true,
					ContributingGoals: []string{"BS Mathematics"}},
				{Name: "IU Modern Algebra", Category: "Mathematics", Active: true, Period: period,
					ContributingGoals: []string{}, SubProjects: []*planning.Project{{Name: "Read book"}}},
				{Name: "Study Analysis Burkin", Category: "Mathematics", Active: true,
					ContributingGoals: []string{}, SubProjects: []*planning.Project{{Name: "Follow another list"}}},
				{Name: "Study Logic for Mathematicians", Category: "Mathematics", Active: false,
					ContributingGoals: []string{}},
			},
		},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		p := NewParser(l)
		items := p.Parse()

		Eval("projects", items)
		projects := planning.GetRepository().ListProjects()

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
			if !equalGoals(projects[i].ContributingGoals, p.ContributingGoals) {
				t.Errorf("Project has wrong goals. Got %v, want %v", projects[i].ContributingGoals, p.ContributingGoals)
			}
			if !equalProjects(projects[i].SubProjects, p.SubProjects) {
				t.Errorf("Project has wrong subprojects. Got %v, want %v", projects[i].SubProjects, p.SubProjects)
			}
		}
	}
}

func equalProjects(first []*planning.Project, second []*planning.Project) bool {
	if len(first) != len(second) {
		return false
	}
	for i, _ := range first {
		if !(first[i].Name == second[i].Name) {
			return false
		}
	}
	return true
}

func equalGoals(first []string, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	for i, _ := range first {
		if !(first[i] == second[i]) {
			return false
		}
	}
	return true
}
