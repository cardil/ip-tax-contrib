package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/yaml"
)

func Load(file string) (Config, error) {
	log.
		WithField("configPath", file).
		Debug("Loading config as YAML")
	var cfg Config
	bytes, err := os.ReadFile(file)
	if err != nil {
		return Config{}, err
	}
	err = yaml.Unmarshal(bytes, &cfg)
	return cfg, err
}
