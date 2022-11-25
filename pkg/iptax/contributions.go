package iptax

import (
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/cardil/ip-tax-contributions/pkg/config"
)

type Site struct {
	Address string
	config.Type
}

type Repo struct {
	Owner string
	Name  string
	Site
}

func (r Repo) Ref() string {
	sb := strings.Builder{}
	sb.WriteString(r.Owner)
	sb.WriteString("/")
	sb.WriteString(r.Name)
	return sb.String()
}

func (r Repo) URL() url.URL {
	u, _ := url.Parse(r.Site.Address)
	u.Path = path.Join(u.Path, r.Owner, r.Name)
	return *u
}

type Contributions []Contribution

func (c Contributions) ToRepos() []Repo {
	repos := make([]Repo, 0, len(c))
	for _, cc := range c {
		if contains(repos, cc.Repo) {
			continue
		}
		repos = append(repos, cc.Repo)
	}
	return repos
}

func contains(repos []Repo, repo Repo) bool {
	for _, r := range repos {
		if r == repo {
			return true
		}
	}
	return false
}

type Contribution struct {
	Title  string
	Number uint
	Repo
}

func (c Contribution) Ref() string {
	sb := strings.Builder{}
	sb.WriteString(c.Repo.Owner)
	sb.WriteString("/")
	sb.WriteString(c.Repo.Name)
	if c.Site.Type == config.TypeGitLab {
		sb.WriteString("!")
	} else {
		sb.WriteString("#")
	}
	sb.WriteString(strconv.Itoa(int(c.Number)))
	return sb.String()
}

func (c Contribution) URL() url.URL {
	u, _ := url.Parse(c.Site.Address)
	u.Path = path.Join(u.Path, c.Repo.Owner, c.Repo.Name, "pulls", strconv.Itoa(int(c.Number)))
	return *u
}

func CollectContributions(period Period, cfg config.Config) (Contributions, error) {
	contributions := make(Contributions, 0)
	for _, site := range cfg.Sites {

	}
	return contributions, nil
}
