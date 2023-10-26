package planning

import (
	"github.com/bymarkone/timelanguage/internal/utils"
)

type Project struct {
	Id                string
	Name              string
	Category          string
	Period            utils.Period
	ContributingGoals []string
	Active            bool
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
	return FlattenProjectsDepth(arr)
}

func FlattenProjectsDepth(arr []*Project) []*Project {
	var results []*Project
	for i := range arr {
		results = append(results, arr[i])
	}
	return results
}

var ProjectActive = func(val Project) bool {
	return val.Active == true
}
