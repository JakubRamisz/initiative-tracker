package model

type InitiativeObject interface {
	GetName() string
	GetInitiative() int
	CompareInitiative(InitiativeObject) int
}
