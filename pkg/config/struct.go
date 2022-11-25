package config

type Config struct {
	Sites []Site `json:"sites"`
}

type Type string

const (
	TypeGitHub Type = "github"
	TypeGitLab Type = "gitlab"
)

type Site struct {
	Type     `json:"type"`
	Address  string `json:"address"`
	Auth     `json:"auth"`
	Includes []RepoCoordicate `json:"includes"`
	Excludes []RepoCoordicate `json:"excludes"`
}

type Auth struct {
	Token string `json:"token"`
}

type RepoCoordicate struct {
	Org  string `json:"org"`
	Repo string `json:"repo"`
}
