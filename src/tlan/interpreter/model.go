package interpreter

type Persistent interface {
	Persist()
}

type Project struct {
	Name              string
	Category          string
	Start             *Day
	End               *Day
	ContributingGoals []*Goal
	Tasks             []*Task
}

type Goal struct {
	Description string
}

type Task struct {
	Description string
}

type Track struct {
	Name string
}

type Schedule struct {
	Name string
}

type Day struct {
	Day   int
	Month int
	Year  int
}

type Time struct {
	Hour   int
	Minute int
}

type Period struct {
	Start *Time
	End   *Time
}
