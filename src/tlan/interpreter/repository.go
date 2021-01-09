package interpreter

var projects []*Project

func AddProject(project Project)  {
	projects = append(projects, &project)
}

func ListProjects() []*Project {
	return projects
}
