package plan

var projects []*Project

func AddProject(project *Project) {
	projects = append(projects, project)
}

func ListProjects() []*Project {
	return projects
}

func ListProjectsFiltered(cond func(project Project) bool) []*Project {
	return FilterProjects(projects, cond)
}

func Clean() {
	projects = []*Project{}
}
