package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFilename = "/.gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	err := c.write()
	if err != nil {
		return err
	}
	return nil
}

func (c Config) write() error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(c)
	if err != nil {
		return err
	}
	os.WriteFile(configPath, jsonData, os.ModeAppend)
	fmt.Println("ok")
	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home + configFilename, nil
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}

	defer file.Close()
	data := make([]byte, 256)
	// Read file, store byte slice in *data* and bytecount in *n*
	n, err := file.Read(data)
	if err != nil {
		return Config{}, err
	}
	var config Config
	//
	err = json.Unmarshal(data[:n], &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
