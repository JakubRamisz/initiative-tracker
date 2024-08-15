package test

import (
	"itr/logic"
	"os"
	"slices"
	"testing"
)

func TestReadConfig(t *testing.T) {
	config, err := logic.LoadFromFile("test_conf.json")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("config: %+v", config)
}

func TestWriteConfig(t *testing.T) {
	const path = "test_write_conf"

	os.Remove(path)

	config := logic.Config{
		PlayerCharacters: []string{"Jaturn", "Kier"},
	}

	err := config.SaveToFile(path)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)

	exists, err := logic.CheckFileExists(path)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("file does not exist")
	}

	savedConf, err := logic.LoadFromFile(path)
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
	exists, err := logic.CheckFileExists("test_conf.json")
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatalf("expected: %t got: %t", true, exists)
	}

	exists, err = logic.CheckFileExists("manbearpig")
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatalf("expected: %t got: %t", false, exists)
	}
}
