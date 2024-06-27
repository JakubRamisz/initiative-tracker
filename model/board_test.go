package model

import (
	"fmt"
	"testing"
)

func TestOrderInitiative(t *testing.T) {
	npcs := []NPC{
		{
			Creature: Creature{
				Name:       "npc17",
				Initiative: 17,
			},
		},
		{
			Creature: Creature{
				Name:       "npc20",
				Initiative: 20,
			},
		},
	}
	pcs :=
		[]Creature{
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
		[]Event{
			{
				Name:       "ev20",
				Initiative: 20,
			},
		}

	b := Board{}
	for _, pc := range pcs {
		pc := pc
		b.AddPC(&pc)
	}
	for _, npc := range npcs {
		npc := npc
		b.AddNPC(&npc)
	}
	for _, event := range events {
		event := event
		b.AddEvent(&event)
	}
	objects := b.GetObjects()
	for i, o := range objects {
		fmt.Printf("%d %s \n", o.GetInitiative(), o.GetName())
		if i+1 == len(objects) {
			return
		}

		if objects[i].GetInitiative() <
			objects[i+1].GetInitiative() {
			t.Fatalf("initiative ordered incorrectly")
		}
		if (objects[i].GetInitiative() == objects[i+1].GetInitiative()) && (objects[i].GetPriority() < objects[i+1].GetPriority()) {
			t.Fatalf("priority ordered incorrectly")
		}
	}
}
