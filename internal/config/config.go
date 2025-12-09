package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".gatorconfig.json"), nil
}

func readConfigFromFile(path string) (Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	var cfg Config
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	cfg, err := readConfigFromFile(path)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func writeConfigToFile(path string, cfg Config) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(cfg)
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	return writeConfigToFile(path, *c)
}
