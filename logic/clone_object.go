package logic

import (
	"errors"
	"itr/models"
	"slices"
)

func (b *Board) CloneObject(obj models.BoardObject) error {
	if obj == nil {
		return errors.New("object is nil")
	}
	if !slices.Contains(b.objects, obj) {
		return errors.New("object not found")
	}

	if _, ok := obj.(*models.Creature); ok {
		return errors.New("cannot clone player characters")
	}

	var clone models.BoardObject
	if npc, ok := obj.(*models.NPC); ok {
		exists := 0
		for _, obj := range b.objects {
			if obj.GetName() == npc.Name {
				exists++
			}
		}
		clone = &models.NPC{
			Creature: models.Creature{
				Name: npc.Name,
			},
			MaxHP:          npc.MaxHP,
			CurrentHP:      npc.CurrentHP,
			CreatureNumber: exists,
		}
	} else if ev, ok := obj.(*models.Event); ok {
		clone = &models.Event{
			Name: ev.Name,
		}
	}

	if clone == nil {
		return errors.New("unknown object type")
	}
	if err := b.add(clone); err != nil {
		return err
	}
	if err := b.SetInitiative(clone, obj.GetInitiative()); err != nil {
		return err
	}
	return nil
}
