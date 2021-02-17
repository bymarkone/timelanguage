package planning

import (
	"tlan/utils"
)

type Project struct {
	Name              string
	Category          string
	Type              string
	Period            utils.Period
	ContributingGoals []string
	SubProjects       []*Project
	Parent            *Project
	Active            bool
	Level             int
}

func (p *Project) AllSubProjects() []*Project {
	var result []*Project
	for _, subProject := range p.SubProjects {
		result = append(result, subProject)
		result = append(result, subProject.AllSubProjects()...)
	}
	return result
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
		results = append(results, FlattenProjectsDepth(arr[i].SubProjects)...)
	}
	return results
}

var ByActive = func(val Project) bool {
	return val.Active == true
}

var ByInactive = func(val Project) bool {
	return val.Active == false
}
