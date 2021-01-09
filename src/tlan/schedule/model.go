package plan

import "tlan/plan"

type Track struct {
	Name string
	Project []*plan.Project
}

type Schedule struct {
	Name string
}

type Time struct {
	Hour   int
	Minute int
}

type Period struct {
	Start *Time
	End   *Time
}
