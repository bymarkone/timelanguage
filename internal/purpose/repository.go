package purpose

const GoalLess = "GoalLess"

var goalLess *Goal

func init() {
	goalLess = &Goal{Name: GoalLess, Category: GoalLess}
}

type Repository struct {
	goals  []*Goal
	values []*Value
}

var repository Repository

func CreateRepository() {
	repository = Repository{}
}

func GetRepository() *Repository {
	return &repository
}

func (r *Repository) AddGoal(goal *Goal) {
	r.goals = append(r.goals, goal)
}

func (r *Repository) ListGoals() []*Goal {
	return r.goals
}

func (r *Repository) GetGoal(name string) *Goal {
	if name == GoalLess {
		return goalLess
	}
	return FindGoal(r.goals, ByGoalName(name))
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

func GoalsByCategory() map[string][]*Goal {
	var result = make(map[string][]*Goal)
	for _, goal := range repository.goals {
		result[goal.Category] = append(result[goal.Category], goal)
	}
	return result
}

func (r *Repository) ListValues() []*Value {
	return r.values
}

func (r *Repository) AddValue(value *Value) {
	r.values = append(r.values, value)
}
