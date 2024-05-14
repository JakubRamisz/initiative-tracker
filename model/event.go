package model

type Event struct {
	Name       string
	Initiative int
}

func (e Event) CompareInitiative(other InitiativeObject) int {
	diff := e.Initiative - other.GetInitiative()
	if diff > 0 {
		return 1
	}
	return -1

}

func (e Event) GetName() string {
	return e.Name
}

func (e Event) GetInitiative() int {
	return e.Initiative
}
