package logic

import "itr/models"

func (b *Board) AddEvent(name string, initiative int) error {
	event := models.Event{
		Name:       name,
		Initiative: initiative,
	}
	b.add(&event)

	return nil
}
