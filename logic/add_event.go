package logic

import "itr/models"

func (b *Board) AddEvent(name string, initiative int) error {
	event := models.Event{
		Name:       name,
		Initiative: initiative,
	}
	if err := b.add(&event); err != nil {
		return err
	}

	return nil
}
