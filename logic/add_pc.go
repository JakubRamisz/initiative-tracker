package logic

import (
	"errors"
	"itr/models"
	"slices"
)

func (b *Board) AddPC(name string) error {
	exists := false
	for _, obj := range b.objects {
		if obj.GetName() == name {
			exists = true
			break
		}
	}
	if exists {
		err := errors.New("a player character with that name already exists")
		return err
	}
	pc := models.Creature{
		Name: name,
	}
	if err := b.add(&pc); err != nil {
		return err
	}

	conf, err := LoadFromFile(ConfigFilePath)
	if err != nil {
		return err
	}
	if slices.Contains(conf.PlayerCharacters, pc.Name) {
		return nil
	}
	conf.PlayerCharacters = append(conf.PlayerCharacters, pc.Name)

	err = conf.SaveToFile(ConfigFilePath)
	if err != nil {
		return err
	}
	return nil
}
