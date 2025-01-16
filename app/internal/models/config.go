package models

import (
	"fmt"

	"GO.CONSOLE.TODO/internal/storage"
)

type Config struct {
	AppName    string
	Version    string
	WelcomeMsg string
}

const (
	DEFAULT_CONFIG_FILE = "jsonConfig.json"
)

func GetConfig(filename string) (*Config, error) {
	result, err := storage.ReadJson[Config](filename)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}
	return result, nil
}
