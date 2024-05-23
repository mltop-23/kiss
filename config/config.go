package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config struct holds configuration settings
type Config struct {
	Port     int `json:"port"`
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"database"`
}

// LoadConfig reads configuration from a JSON file
func LoadConfig() (*Config, error) {
	// Replace "config.json" with your actual file name
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return &config, nil
}
