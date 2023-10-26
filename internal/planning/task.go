package planning

type Task struct {
	Active  bool
	Urgent  bool
	Name    string
	Type    string
	Project Project
}

func FilterTasks(arr []*Task, cond func(Task) bool) []*Task {
	var result []*Task
	for i := range arr {
		if cond(*arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}

var ByType = func(taskType string) func(val Task) bool {
	return func(val Task) bool {
		return val.Type == taskType
	}
}

var TaskActive = func(val Task) bool {
	return val.Active == true
}
