package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
)

type Configuration struct {
	// Number of tags to keep
	TagsToKeep int
	// Default container registry
	Registry string
}

func NewConfig() (*Configuration, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	return mergeWithDefaultsReflect(config), nil
}

func mergeWithDefaultsReflect(config *Configuration) *Configuration {
	defaults := getDefaultConfig()
	if config == nil {
		return defaults
	}

	result := *defaults
	rValue := reflect.ValueOf(config).Elem()
	rResult := reflect.ValueOf(&result).Elem()

	for i := 0; i < rValue.NumField(); i++ {
		field := rValue.Field(i)
		switch field.Kind() {
		case reflect.String:
			if field.String() != "" {
				rResult.Field(i).SetString(field.String())
			}
		case reflect.Int:
			if field.Int() != 0 {
				rResult.Field(i).SetInt(field.Int())
			}
		}
	}

	return &result
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

func getDefaultConfig() *Configuration {
	config := &Configuration{
		TagsToKeep: 10,
		Registry:   "",
	}

	return config
}

func createDefaultConfig() (*Configuration, error) {
	config := getDefaultConfig()

	if err := SaveConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}
