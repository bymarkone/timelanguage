package purpose

import plan "tlan/planning"

type Goal struct {
	Category string
	Name     string
	Tags     []string
	Projects []*plan.Project
}
