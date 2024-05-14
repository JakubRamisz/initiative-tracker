package model

type Game struct {
	PCs          []PC
	NPCs         []NPC
	Events       []Event
	CurrentRound int
}
