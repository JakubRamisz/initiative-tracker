package cmd

import (
	"itr/logic"
	"itr/model"
	"os"
	"slices"
	"testing"
)

func TestRemovePC(t *testing.T) {
	const path = "../test/test_remove_pcs"
	const pc = "Illya"

	conf := model.Config{
		PlayerCharacters: []string{pc},
	}

	err := logic.WriteConfig(conf, path)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)

	err = removePCs([]string{pc}, path)
	if err != nil {
		t.Fatal(err)
	}
	savedConf, err := logic.ReadConfig(path)
	if err != nil {
		t.Fatal(err)
	}

	if slices.Contains(savedConf.PlayerCharacters, pc) {
		t.Fatal("failed to remove a pc")
	}
}
