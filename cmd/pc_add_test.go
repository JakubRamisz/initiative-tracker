package cmd

import (
	"itr/logic"
	"itr/model"
	"os"
	"slices"
	"testing"
)

func TestAddPC(t *testing.T) {
	const path = "../test/test_add_pcs"
	const pc = "Illya"

	conf := model.Config{
		PlayerCharacters: []string{pc},
	}

	err := logic.WriteConfig(conf, path)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)

	err = addPCs([]string{pc}, path)
	if err != nil {
		t.Fatal(err)
	}
	savedConf, err := logic.ReadConfig(path)
	if err != nil {
		t.Fatal(err)
	}

	if !slices.Contains(savedConf.PlayerCharacters, pc) {
		t.Fatal("failed to add a pc")
	}
}
