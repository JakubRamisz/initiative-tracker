package models

type BoardObject interface {
	GetInfo() string
	GetName() string
	GetInitiative() int
	GetPriority() int
	CompareInitiative(BoardObject) int
	SetInitiative(int)
	SetPriority(int)
}

type BoardObjects []BoardObject

func (g BoardObjects) Len() int {
	return len(g)
}

func (g BoardObjects) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g BoardObjects) Less(i, j int) bool {
	return g[i].CompareInitiative(g[j]) > 0
}
