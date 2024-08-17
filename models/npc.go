package models

import "fmt"

type NPC struct {
	Creature
	CurrentHP      int
	MaxHP          int
	CreatureNumber int
}

func (npc *NPC) Heal(value int) {
	newValue := npc.CurrentHP + value
	if newValue <= npc.MaxHP {
		npc.CurrentHP = newValue
	} else {
		npc.CurrentHP = npc.MaxHP
	}
}

func (npc *NPC) RecieveDMG(value int) {
	newValue := npc.CurrentHP - value
	if newValue >= 0 {
		npc.CurrentHP = newValue
	} else {
		npc.CurrentHP = 0
	}
}

func (npc *NPC) GetInfo() string {
	name := npc.Name
	if npc.CreatureNumber > 0 {
		name = fmt.Sprintf("%s (%d)", name, npc.CreatureNumber)
	}
	return fmt.Sprintf("%2d %-20s [%d/%d]", npc.Initiative, name, npc.CurrentHP, npc.MaxHP)
}
