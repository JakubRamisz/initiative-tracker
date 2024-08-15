package logic

import (
	"errors"
	"itr/models"
)

func (b *Board) RemoveObject(obj models.BoardObject) error {
	found := false
	for i, o := range b.objects {
		if o == obj {
			b.objects = append(b.objects[:i], b.objects[i+1:]...)
			found = true
		}
	}
	if !found {
		return errors.New("obj not found")
	}

	if pc, ok := obj.(*models.Creature); ok {
		conf, err := LoadFromFile(ConfigFilePath)
		if err != nil {
			return err
		}
		var newPCs []string
		for _, confPC := range conf.PlayerCharacters {
			if confPC == pc.Name {
				continue
			}

			newPCs = append(newPCs, confPC)
		}
		conf.PlayerCharacters = newPCs

		err = conf.SaveToFile(ConfigFilePath)
		return err
	}
	return nil
}
