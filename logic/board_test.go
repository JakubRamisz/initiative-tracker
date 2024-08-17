package logic

import (
	"fmt"
	"itr/models"
	"testing"
)

func TestOrderInitiative(t *testing.T) {
	npcs := []models.NPC{
		{
			Creature: models.Creature{
				Name:       "npc17",
				Initiative: 17,
			},
		},
		{
			Creature: models.Creature{
				Name:       "npc20",
				Initiative: 20,
			},
		},
	}
	pcs :=
		[]models.Creature{
			{
				Name:       "char20",
				Initiative: 20,
			},
			{
				Name:       "char13 p0",
				Initiative: 13,
				Priority:   0,
			},
			{
				Name:       "char7",
				Initiative: 7,
			},
			{
				Name:       "char13 p1",
				Initiative: 13,
				Priority:   1,
			},
		}
	events :=
		[]models.Event{
			{
				Name:       "ev20",
				Initiative: 20,
			},
		}

	b := Board{}
	for _, pc := range pcs {
		pc := pc
		if err := b.add(&pc); err != nil {
			t.Fatalf("failed to add pc")
		}
	}
	for _, npc := range npcs {
		npc := npc
		if err := b.add(&npc); err != nil {
			t.Fatalf("failed to add npc")
		}
	}
	for _, event := range events {
		event := event
		if err := b.add(&event); err != nil {
			t.Fatalf("failed to add event")
		}
	}
	for i, o := range b.objects {
		fmt.Printf("%d %s \n", o.GetInitiative(), o.GetName())
		if i+1 == len(b.objects) {
			return
		}

		if b.objects[i].GetInitiative() <
			b.objects[i+1].GetInitiative() {
			t.Fatalf("initiative ordered incorrectly")
		}
		if (b.objects[i].GetInitiative() == b.objects[i+1].GetInitiative()) &&
			(b.objects[i].GetPriority() < b.objects[i+1].GetPriority()) {
			t.Fatalf("priority ordered incorrectly")
		}
	}
}
