package logic

import (
	"itr/models"
)

func (b *Board) AddNPC(name string, hp int) error {
	exists := 0
	for _, obj := range b.objects {
		if obj.GetName() == name {
			exists++
		}
	}

	npc := models.NPC{
		Creature: models.Creature{
			Name: name,
		},
		MaxHP:          hp,
		CurrentHP:      hp,
		CreatureNumber: exists,
	}
	if err := b.add(&npc); err != nil {
		return err
	}

	return nil
}
