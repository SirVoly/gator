package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	var cfg Config

	filepath, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}

	dat, err := os.ReadFile(filepath)
	if err != nil {
		return cfg, err
	}

	if err = json.Unmarshal(dat, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func write(cfg Config) error {
	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	dat, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	if err = os.WriteFile(filepath, dat, 0600); err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homePath + string(os.PathSeparator) + configFileName, nil
}
