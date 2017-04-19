package github

import (
	"github.com/xanzy/go-gitlab"
)

// Procjet represents a single project from gitlab.
// This struct is a stripped down version of gitlab.Project.
// We only return the values we need here.
type Project struct {
	StarCount int `json:"star_count"`
}

// GetRepositoryDetails will retrieve details about the repository owner/repo from github.
func GetProjectDetails(pid int) (*Project, error) {
	client := gitlab.NewClient(nil)
	client.SetBaseURL("https://gitlab.com/api/v3")
	project, _, err := client.Projects.GetProject(pid)
	if project == nil {
		return nil, err
	}

	r := &Project{
		StarCount: repository.StarCount,
	}
	return r, err
}
