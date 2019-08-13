package gitlab

import (
	"github.com/xanzy/go-gitlab"
)

type Config struct {
	Token string
	Ref   string
	Pid   int
	Url   string // default `https://gitlab.com/api/v4`
}

type catcher struct {
	config Config
	git    *gitlab.Client
}

func Catcher(c Config) *catcher {
	git := gitlab.NewClient(nil, c.Token)
	if len(c.Url) != 0 {
		git.SetBaseURL(c.Url)
	}
	return &catcher{
		config: c,
		git:    git,
	}
}

func (c *catcher) Fetch(fileName string) ([]byte, error) {
	data, _, err := c.git.RepositoryFiles.GetRawFile(c.config.Pid, fileName, &gitlab.GetRawFileOptions{Ref: &c.config.Ref})
	if err != nil {
		return nil, err
	}
	return data, nil
}
