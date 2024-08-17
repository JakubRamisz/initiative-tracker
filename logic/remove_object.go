package logic

import (
	"errors"
	"itr/models"
	"slices"
)

func (b *Board) RemoveObject(obj models.BoardObject) error {
	if obj == nil {
		return errors.New("object is nil")
	}
	if !slices.Contains(b.objects, obj) {
		return errors.New("object not found")
	}
	for i, o := range b.objects {
		if o == obj {
			b.objects = append(b.objects[:i], b.objects[i+1:]...)
		}
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
