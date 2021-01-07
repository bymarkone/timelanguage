package interpreter

var projects []*Project

func addProject(project Project)  {
	projects = append(projects, &project)
}

func listProjects() []*Project {
	return projects
}