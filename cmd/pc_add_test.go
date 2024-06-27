package cmd

import (
	"itr/utils"
	"os"
	"slices"
	"testing"
)

func TestAddPC(t *testing.T) {
	const path = "../test/test_add_pcs"
	const pc = "Illya"

	conf := utils.Config{
		PlayerCharacters: []string{pc},
	}

	err := conf.SaveToFile(path)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)

	err = addPCs([]string{pc}, path)
	if err != nil {
		t.Fatal(err)
	}
	savedConf, err := utils.LoadFromFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if !slices.Contains(savedConf.PlayerCharacters, pc) {
		t.Fatal("failed to add a pc")
	}
}
