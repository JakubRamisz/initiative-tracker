package models

import "fmt"

type NPC struct {
	Creature
	CurrentHP int
	MaxHP     int
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
	return fmt.Sprintf("%2d %-20s [%d/%d]", npc.Initiative, npc.Name, npc.CurrentHP, npc.MaxHP)
}
