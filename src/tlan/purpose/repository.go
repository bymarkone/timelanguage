package purpose

var goals []*Goal

const GoalLess = "GoalLess"
var goalLess *Goal

func init() {
	goalLess  = &Goal{Name: GoalLess, Category: GoalLess}
}

func AddGoal(goal *Goal) {
	goals = append(goals, goal)
}

func ListGoals() []*Goal {
	return goals
}

func GetGoal(name string) *Goal {
	if name == GoalLess {
		return goalLess
	}
	return FindGoal(goals, ByGoalName(name))
}

func ByGoalName(name string) func(goal Goal) bool {
	return func(goal Goal) bool { return goal.Name == name }
}

func FindGoal(arr []*Goal, cond func(goal Goal) bool) *Goal {
	var result *Goal
	for i := range arr {
		if cond(*arr[i]) {
			result = arr[i]
		}
	}
	return result
}
