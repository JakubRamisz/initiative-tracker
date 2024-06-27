package model

type BoardObject interface {
	GetInfo() string
	GetName() string
	GetInitiative() int
	GetPriority() int
	CompareInitiative(BoardObject) int
	SetInitiative(int)
	SetPriority(int)
}
