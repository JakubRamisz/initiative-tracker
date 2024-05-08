package logic

import (
	"itr/model"
	"os"
	"slices"
	"testing"
)

func TestReadConfig(t *testing.T) {
	config, err := ReadConfig("../test/test_conf.json")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("config: %+v", config)
}

func TestWriteConfig(t *testing.T) {
	const path = "../test/test_write_conf"

	os.Remove(path)

	config := model.Config{
		PlayerCharacters: []string{"Jaturn", "Kier"},
	}

	err := WriteConfig(config, path)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)

	exists, err := checkFileExists(path)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("file does not exist")
	}

	savedConf, err := ReadConfig(path)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("config: %+v", config)
	t.Logf("saved config: %+v", savedConf)
	if !slices.Equal(config.PlayerCharacters, savedConf.PlayerCharacters) {
		t.Fatal("config saved incorrectly")
	}

}

func TestCheckFileExists(t *testing.T) {
	exists, err := checkFileExists("../test/test_conf.json")
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatalf("expected: %t got: %t", true, exists)
	}

	exists, err = checkFileExists("manbearpig")
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatalf("expected: %t got: %t", false, exists)
	}
}
