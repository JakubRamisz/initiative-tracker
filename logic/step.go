package logic

import (
	"errors"
	"itr/models"
)

func (b *Board) Step() error {
	if len(b.objects) == 0 {
		return errors.New("no objects")
	}

	currentIdx := -1
	if b.currentIndex != nil {
		currentIdx = *b.currentIndex
	}
	for i, o := range b.objects {
		i := i
		o := o
		if i <= currentIdx {
			continue
		}
		if npc, ok := o.(*models.NPC); ok && npc.CurrentHP <= 0 {
			continue
		}
		b.currentIndex = &i
		return nil
	}

	for i, o := range b.objects {
		i := i
		o := o
		if i >= currentIdx {
			continue
		}
		if npc, ok := o.(*models.NPC); ok && npc.CurrentHP <= 0 {
			continue
		}
		b.currentIndex = &i
		b.round++
		return nil
	}
	b.currentIndex = nil
	return nil
}
