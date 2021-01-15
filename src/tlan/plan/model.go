package plan

import "tlan/utils"

type Project struct {
	Name              string
	Category          string
	Period            utils.Period
	ContributingGoals []*Goal
	Tasks             []*Task
	Active            bool
}

type Goal struct {
	Description string
}

type Task struct {
	Description string
}

func FilterProjects(arr []*Project, cond func(project Project) bool) []*Project {
	var result []*Project
	for i := range arr {
		if cond(*arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}

var ByActive = func(val Project) bool {
	return val.Active == true
}

var ByInactive = func(val Project) bool {
	return val.Active == false
}
