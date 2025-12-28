package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	BranchPrefixes map[string]string `json:"branchPrefixes"`
	BaseBranches   map[string]string `json:"baseBranches"`
}

func DefaultConfig() Config {
	return Config{
		BranchPrefixes: map[string]string{
			"feature": "feature/",
			"hotfix":  "hotfix/",
			"release": "release/",
		},
		BaseBranches: map[string]string{
			"feature": "develop",
			"hotfix":  "main",
			"release": "develop",
		},
	}
}

func LoadConfig() (Config, error) {
	config, err := loadFromPath(".gwarc.json")
	if err == nil {
		return config, nil
	}

	home, err := os.UserHomeDir()
	if err == nil {
		config, err = loadFromPath(filepath.Join(home, ".gwarc.json"))
		if err == nil {
			return config, nil
		}
	}

	return DefaultConfig(), nil
}

func loadFromPath(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func SaveConfig(config Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(".gwarc.json", data, 0644)
}