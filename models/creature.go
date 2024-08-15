package models

import "fmt"

type Creature struct {
	Name       string
	Initiative int
	Priority   int
}

func (c *Creature) CompareInitiative(other BoardObject) int {
	diff := c.Initiative - other.GetInitiative()
	if diff > 0 {
		return 1
	}
	if diff < 0 {
		return -1
	}

	priorityDiff := c.Priority - other.GetPriority()
	if priorityDiff > 0 {
		return 1
	}
	if priorityDiff < 0 {
		return -1
	}
	return 0
}

func (c *Creature) GetPriority() int {
	return c.Priority
}

func (c *Creature) GetName() string {
	return c.Name
}

func (c *Creature) GetInitiative() int {
	return c.Initiative
}

func (c *Creature) GetInfo() string {
	return fmt.Sprintf("%2d %-20s", c.Initiative, c.Name)
}

func (c *Creature) SetInitiative(initiative int) {
	c.Initiative = initiative
}

func (c *Creature) SetPriority(priority int) {
	c.Priority = priority
}
