package plan

type Project struct {
	Name              string
	Category          string
	Start             Day
	End               Day
	ContributingGoals []*Goal
	Tasks             []*Task
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

type Goal struct {
	Description string
}

type Task struct {
	Description string
}

type Day struct {
	Day   int
	Month int
	Year  int
}

