package logic

import (
	"errors"
	"itr/models"
	"slices"
	"sort"
)

func (b *Board) SetInitiative(obj models.BoardObject, initiative int) error {
	if obj == nil {
		return errors.New("object is nil")
	}
	if !slices.Contains(b.objects, obj) {
		return errors.New("object not found")
	}
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
