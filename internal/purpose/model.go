package purpose

import (
	plan "github.com/bymarkone/timelanguage/internal/planning"
)

type Goal struct {
	Category string
	Name     string
	Tags     []string
	Projects []*plan.Project
}

type Value struct {
	Name string
}
