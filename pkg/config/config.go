package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Configuration struct {
	// Number of tags to keep
	TagsToKeep int
}

func NewConfig() (*Configuration, error) {
	return loadConfig()
}

func getConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, ".config", "azutils")
	return configDir, nil
}

func getConfigPath() (string, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "config.json"), nil
}
func loadConfig() (*Configuration, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	// If config doesn't exist, create default config
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return createDefaultConfig()
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Configuration
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func SaveConfig(config *Configuration) error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

func createDefaultConfig() (*Configuration, error) {
	config := &Configuration{
		TagsToKeep: 10,
	}

	if err := SaveConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}
