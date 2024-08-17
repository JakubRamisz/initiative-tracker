package logic

import (
	"fmt"
	"itr/models"
)

func (b *Board) AddNPC(name string, hp int) error {
	exists := 0
	for _, obj := range b.objects {
		if obj.GetName() == name {
			exists++
		}
	}
	if exists > 0 {
		name = fmt.Sprintf("%s (%d)", name, exists)
	}

	npc := models.NPC{
		Creature: models.Creature{
			Name: name,
		},
		MaxHP:     hp,
		CurrentHP: hp,
	}
	if err := b.add(&npc); err != nil {
		return err
	}

	return nil
}
