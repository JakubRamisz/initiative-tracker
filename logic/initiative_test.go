package logic

import (
	"fmt"
	"itr/model"
	"testing"
)

func TestOrderInitiative(t *testing.T) {
	game := model.Game{
		NPCs: []model.NPC{
			{
				PC: model.PC{
					Name:       "npc17",
					Initiative: 17,
				},
			},
			{
				PC: model.PC{
					Name:       "npc20",
					Initiative: 20,
				},
			},
		},
		PCs: []model.PC{
			{
				Name:       "char20",
				Initiative: 20,
			},
			{
				Name:       "char7",
				Initiative: 7,
			},
			{
				Name:       "char13",
				Initiative: 13,
			},
		},
		Events: []model.Event{
			{
				Name:       "ev20",
				Initiative: 20,
			},
		},
	}

	objects := GetOrderedInitiative(game)
	for i, o := range objects {
		fmt.Printf("%d %s \n", o.GetInitiative(), o.GetName())
		if i+1 == len(objects) {
			return
		}

		if objects[i].GetInitiative() < objects[i+1].GetInitiative() {
			t.Fatalf("initiative ordered incorrectly")
		}
	}
}
