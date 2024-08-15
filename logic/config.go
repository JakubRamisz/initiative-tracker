package logic

import (
	"encoding/json"
	"os"
)

type Config struct {
	PlayerCharacters []string
}

func LoadFromFile(path string) (Config, error) {
	config := Config{}

	fileExists, err := CheckFileExists(path)
	if err != nil {
		return config, err
	}
	if !fileExists {
		return config, nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func (c *Config) SaveToFile(path string) error {
	content, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func CheckFileExists(path string) (bool, error) {
	_, err := os.Stat(path)

	isNotExists := os.IsNotExist(err)
	if err != nil && !isNotExists {
		return false, err
	}

	if isNotExists {
		return false, nil
	}

	return true, nil
}
