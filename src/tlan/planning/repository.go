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

func (r *Repository) ListProjectsFiltered(cond func(project Project) bool) []*Project {
	return FilterProjects(r.projects, cond)
}
