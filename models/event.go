package models

import "fmt"

type Event struct {
	Name       string
	Initiative int
	Priority   int
}

func (e *Event) CompareInitiative(other BoardObject) int {
	diff := e.Initiative - other.GetInitiative()
	if diff > 0 {
		return 1
	}
	if diff < 0 {
		return -1
	}
	priorityDiff := e.Priority - other.GetPriority()
	if priorityDiff > 0 {
		return 1
	}
	return -1

}

func (e *Event) GetPriority() int {
	return e.Priority
}

func (e *Event) GetName() string {
	return e.Name
}

func (e *Event) GetInitiative() int {
	return e.Initiative
}

func (e *Event) GetInfo() string {
	return fmt.Sprintf("%2d %-20s", e.Initiative, e.Name)
}

func (e *Event) SetInitiative(initiative int) {
	e.Initiative = initiative
}

func (e *Event) SetPriority(priority int) {
	e.Priority = priority
}
