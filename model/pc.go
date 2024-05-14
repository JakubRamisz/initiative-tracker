package model

type PC struct {
	Name       string
	Initiative int
}

func (pc PC) CompareInitiative(other InitiativeObject) int {
	diff := pc.Initiative - other.GetInitiative()
	if diff > 0 {
		return 1
	}
	if diff < 0 {
		return -1
	}

	return 0
}

func (pc PC) GetName() string {
	return pc.Name
}

func (pc PC) GetInitiative() int {
	return pc.Initiative
}
