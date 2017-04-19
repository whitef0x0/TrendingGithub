package trending

import (
	"errors"
	"math/rand"
	"github.com/whitef0x0/go-trending"
)

// TrendingAPI represents the interface to the github.com/trending website
type TrendingAPI interface {
	GetProjects(numberOfPages int) ([]trending.Project, error)
}

// Trend is the data structure to hold a github-trending client.
// This will be used to retrieve trending projects
type Trend struct {
	Client TrendingAPI
}

// NewClient will provide a new instance of Trend.
func NewClient() *Trend {
	githubTrending := trending.NewTrending()

	t := &Trend{
		Client: githubTrending,
	}

	return t
}


// GetRandomProjectGenerator returns a closure to retrieve a random project
func (t *Trend) GetRandomProjectGenerator(numPages int) func() (trending.Project, error) {
	var projects []trending.Project
	var err error

	// Get 1000 most recent trending projects
	projects, err = t.Client.GetProjects(numPages)
	if err != nil {
		return func() (trending.Project, error) {
			return trending.Project{}, err
		}
	}

	// Once we got the projects we will provide a closure
	// to retrieve random projects of this project list.
	return func() (trending.Project, error) {

		// Check the number of projects left in the list
		// If there are no more projects anymore, we will return an error.
		numOfProjects := len(projects)
		if numOfProjects == 0 {
			return trending.Project{}, errors.New("No projects found")
		}

		// If there are projects left, chose a random one ...
		randomNumber := rand.Intn(numOfProjects)
		randomProject := projects[randomNumber]

		// ... and delete the chosen project from our list.
		projects = append(projects[:randomNumber], projects[randomNumber+1:]...)

		return randomProject, nil
	}
}
