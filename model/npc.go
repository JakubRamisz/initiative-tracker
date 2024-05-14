package model

type NPC struct {
	PC
	CurrentHP int
	MaxHP     int
}

func (npc *NPC) RecieveDMG(value int) {
	newValue := npc.CurrentHP - value
	if newValue >= 0 {
		npc.CurrentHP = newValue
	} else {
		npc.CurrentHP = 0
	}
}
