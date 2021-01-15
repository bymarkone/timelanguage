package plan

import "tlan/utils"

type Project struct {
	Name              string
	Category          string
	Period            utils.Period
	ContributingGoals []*Goal
	SubProjects       []*Project
	Active            bool
	Level 						int
}

type Goal struct {
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

func FlattenProjects(arr []*Project) []*Project {
	return flattenProjects(arr)
}

func flattenProjects(arr []*Project) []*Project {
	var results []*Project
	for i := range arr {
		results = append(results, arr[i])
		results = append(results, flattenProjects(arr[i].SubProjects)...)
	}
	return results
}

var ByActive = func(val Project) bool {
	return val.Active == true
}

var ByInactive = func(val Project) bool {
	return val.Active == false
}
