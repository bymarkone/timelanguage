package planning

var repository Repository

type Repository struct {
	projects []*Project
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

func ByProjectName(name string) func(project Project) bool {
	return func(project Project) bool { return project.Name == name }
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
