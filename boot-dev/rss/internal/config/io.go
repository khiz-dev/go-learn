package config

import (
	"encoding/json"
	"os"
)

func ReadConfig(c *Config) error {
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

func WriteConfig(c *Config) error {
	file, err := os.Create("../../.gatorconfig.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		return err
	}

	return nil
}