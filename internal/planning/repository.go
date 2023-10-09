package planning

import (
	"strings"
)

var repository Repository

type Repository struct {
	projects []*Project
	tasks    []*Task
}

func CreateRepository() {
	repository = Repository{}
}

func GetRepository() *Repository {
	return &repository
}

func (r *Repository) AddProject(project *Project) {
	r.projects = append(r.projects, project)
}

func (r *Repository) ListProjects() []*Project {
	return r.projects
}

func (r *Repository) ProjectsByCategory() map[string][]*Project {
	var result = make(map[string][]*Project)
	for _, project := range repository.projects {
		result[project.Category] = append(result[project.Category], project)
	}
	return result
}

func (r *Repository) ListCategories() []string {
	keys := make(map[string]bool)
	var result []string
	for _, project := range r.ListProjects() {
		if _, value := keys[project.Category]; !value {
			keys[project.Category] = true
			result = append(result, project.Category)
		}
	}
	return result
}

func (r *Repository) ListProjectsFiltered(cond func(project Project) bool) []*Project {
	return FilterProjects(r.projects, cond)
}

func (r *Repository) GetProject(name string) *Project {
	return FindProject(r.projects, ByProjectName(name))
}

func (r *Repository) GetProjectById(id string) *Project {
	return FindProject(r.projects, ByProjectId(id))
}

func (r *Repository) GetProjectsWithText(id string) *Project {
	for _, project := range r.ListProjects() {
		if strings.Contains(project.Name, id) {
			return project
		}
	}
	return nil
}

func ByProjectName(name string) func(project Project) bool {
	return func(project Project) bool { return project.Name == name }
}

func ByProjectId(id string) func(project Project) bool {
	return func(project Project) bool { return project.Id == id }
}

func FindProject(arr []*Project, cond func(project Project) bool) *Project {
	var result *Project
	for i := range arr {
		if cond(*arr[i]) {
			result = arr[i]
		}
	}
	return result
}

func (r *Repository) AddTask(task *Task) {
	r.tasks = append(r.tasks, task)
}

func (r *Repository) ListTasks() []*Task {
	return r.tasks
}

type void struct{}

func (r *Repository) TypesFromTasks() []string {
	var member void
	set := make(map[string]void)
	for _, task := range r.tasks {
		set[task.Type] = member
	}
	keys := make([]string, len(set))
	i := 0
	for k := range set {
		keys[i] = k
		i++
	}
	return keys
}
