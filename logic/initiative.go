package logic

import (
	"itr/model"
	"slices"
)

func GetOrderedInitiative(g model.Game) []model.InitiativeObject {
	objects := []model.InitiativeObject{}
	for i := range g.PCs {
		objects = append(objects, g.PCs[i])
	}
	for i := range g.NPCs {
		objects = append(objects, g.NPCs[i])
	}
	for i := range g.Events {
		objects = append(objects, g.Events[i])
	}

	sortFunc := func(val1 model.InitiativeObject, val2 model.InitiativeObject) int {
		return val2.CompareInitiative(val1)
	}

	slices.SortFunc(objects, sortFunc)

	return objects
}
