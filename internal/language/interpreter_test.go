package language

import (
	"github.com/bymarkone/timelanguage/internal/planning"
	"github.com/bymarkone/timelanguage/internal/purpose"
	"github.com/bymarkone/timelanguage/internal/schedule"
	"github.com/bymarkone/timelanguage/internal/utils"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestEvalGoals(t *testing.T) {
	tests := []struct {
		input    string
		expected []*purpose.Goal
	}{
		{
			`
Great Goal [Lagging]
Great Technologist [Lagging]
- Drive strategy and execution [Leading]
- First class engineer [Leading]
- Democratize best practices [Leading]
`,
			[]*purpose.Goal{
				{Name: "Great Goal", Tags: []string{"Lagging"}, Category: "Great Goal"},
				{Name: "Drive strategy and execution", Tags: []string{"Leading"}, Category: "Great Technologist"},
				{Name: "First class engineer", Tags: []string{"Leading"}, Category: "Great Technologist"},
				{Name: "Democratize best practices", Tags: []string{"Leading"}, Category: "Great Technologist"},
			},
		},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		p := NewParser("test_goals", l)
		categories, items := p.Parse()

		Eval("goals", categories, items)
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
Creative Work [Daily. 05:00-09:00]
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
		p := NewParser("test_schedule", l)
		categories, items := p.Parse()

		Eval("schedule", categories, items)
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

func TestEvalTasks(t *testing.T) {

	tests := []struct {
		projectInput string
		taskInput    string
		expected     []*planning.Task
	}{{
		`
AI
- Elements of Statistical Learning [Elements.]
- PhD in AI [PhD.]
`,
		`
Elements
* !Take urgent notes
- Read chapters 4 and 5

PhD
* Apply to TUM
`, []*planning.Task{
			{Name: "Take urgent notes", Type: "None", Project: planning.Project{Name: "Elements of Statistical Learning"}},
			{Name: "Read chapters 4 and 5", Type: "None", Project: planning.Project{Name: "Elements of Statistical Learning"}},
			{Name: "Apply to TUM", Type: "None", Project: planning.Project{Name: "PhD in AI"}}}}}

	for _, tt := range tests {
		l1 := NewLexer(tt.projectInput)
		p1 := NewParser("test_tasks_projects", l1)
		categories1, items1 := p1.Parse()
		Eval("projects", categories1, items1)

		l2 := NewLexer(tt.taskInput)
		p2 := NewParser("test_tasks", l2)
		categories2, items2 := p2.Parse()
		Eval("tasks", categories2, items2)

		tasks := planning.GetRepository().ListTasks()

		for i, p := range tt.expected {
			if tasks[i].Name != p.Name {
				assert.Equal(t, tasks[i].Name, p.Name)
				assert.Equal(t, tasks[i].Project.Name, p.Project.Name)
			}
		}

		assert.True(t, tasks[0].Active, "Task bullet * should me classified as active")
		assert.Falsef(t, tasks[1].Active, "Task bullet ! should me classified as non active")
		assert.True(t, tasks[2].Active, "Task bullet * should me classified as active")
		assert.True(t, tasks[0].Urgent, "Task starting with ! not market as urgent")

	}
}

func TestEvalValues(t *testing.T) {

	tests := []struct {
		input    string
		expected []*purpose.Value
	}{{`
Values
- Family
- Health
`, []*purpose.Value{{Name: "Family"}, {Name: "Health"}}}}

	for _, tt := range tests {
		purpose.CreateRepository()

		l := NewLexer(tt.input)
		p := NewParser("test_values", l)
		categories, items := p.Parse()

		Eval("values", categories, items)

		values := purpose.GetRepository().ListValues()

		for i, p := range tt.expected {
			if values[i].Name != p.Name {
				t.Errorf("Project has wrong data. Got %s, want %s", values[i].Name, p.Name)
			}
		}
	}

}

func TestEvalProjects(t *testing.T) {

	period := utils.Period{Start: utils.DateTime{Day: 10, Month: 1, Year: 2021}, End: utils.DateTime{Day: 15, Month: 4, Year: 2022}}

	tests := []struct {
		input    string
		expected []*planning.Project
	}{
		{
			`
Mathematics
- IU Analysis II >> BS Mathematics
- IU Modern Algebra [Algebra. 10/01/21-15/04/22]
- Study Analysis Burkin
- (Study Logic for Mathematicians)
`,
			[]*planning.Project{
				{Name: "IU Analysis II", Category: "Mathematics", Active: true,
					ContributingGoals: []string{"BS Mathematics"}},
				{Id: "Algebra", Name: "IU Modern Algebra", Category: "Mathematics", Active: true, Period: period,
					ContributingGoals: []string{}},
				{Name: "Study Analysis Burkin", Category: "Mathematics", Active: true,
					ContributingGoals: []string{}},
				{Name: "Study Logic for Mathematicians", Category: "Mathematics", Active: false,
					ContributingGoals: []string{}},
			},
		},
	}

	for _, tt := range tests {
		planning.CreateRepository()

		l := NewLexer(tt.input)
		p := NewParser("test_projects", l)
		categories, items := p.Parse()

		Eval("projects", categories, items)
		projects := planning.GetRepository().ListProjects()

		for i, p := range tt.expected {
			if projects[i].Id != p.Id {
				t.Errorf("Project has wrong data. Got %s, want %s", projects[i].Id, p.Id)
			}
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
			if projects[i].Period.End != p.Period.End {
				t.Errorf("Project has wrong attribute. Got %v, want %v", projects[i].Period.End, p.Period.End)
			}
			if !equalGoals(projects[i].ContributingGoals, p.ContributingGoals) {
				t.Errorf("Project has wrong goals. Got %v, want %v", projects[i].ContributingGoals, p.ContributingGoals)
			}
		}
	}
}

func equalGoals(first []string, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if !(first[i] == second[i]) {
			return false
		}
	}
	return true
}
