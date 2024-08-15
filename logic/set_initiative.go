package logic

import (
	"itr/models"
	"sort"
)

func (b *Board) SetInitiative(obj models.BoardObject, initiative int) error {
	maxPriority := 0
	for _, o := range b.objects {
		if o == obj {
			continue
		}
		if o.GetPriority() > maxPriority {
			maxPriority = o.GetPriority()
		}
	}

	obj.SetInitiative(initiative)
	obj.SetPriority(maxPriority + 1)
	sort.Sort(models.BoardObjects(b.objects))

	return nil
}
