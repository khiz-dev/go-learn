package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBUrl       string `json:"db_url"`
	CurrentUser string `json:"current_user"`
}

func (c *Config) Read() error {
	// Read the configuration from the .gatorconfig.json file
	file, err := os.Open("../../.gatorconfig.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		return err
	}

	return nil
} 