package logic

import (
	"encoding/json"
	"itr/model"
	"os"
)

func ReadConfig(path string) (model.Config, error) {
	config := model.Config{}

	fileExists, err := checkFileExists(path)
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

func WriteConfig(config model.Config, path string) error {
	content, err := json.MarshalIndent(config, "", "\t")
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

func checkFileExists(path string) (bool, error) {
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
