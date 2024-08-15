package logic

import (
	"errors"
	"itr/models"
	"sort"
)

type Board struct {
	objects       []models.BoardObject
	currentIndex  *int
	selectedIndex *int
	round         int
}

func NewBoard(pcNames []string) Board {
	board := Board{}
	for _, pc := range pcNames {
		board.add(
			&models.Creature{
				Name: pc,
			},
		)
	}
	return board
}

func (b *Board) add(n models.BoardObject) {
	b.objects = append(b.objects, n)
	sort.Sort(models.BoardObjects(b.objects))
}

func (b *Board) PCs() (pcs []*models.Creature) {
	for _, o := range b.objects {
		if c, ok := o.(*models.Creature); ok {
			pcs = append(pcs, c)
		}
	}
	return pcs
}

func (b *Board) GetSelectedIndex() *int {
	return b.selectedIndex
}

func (b *Board) GetCurrentIndex() *int {
	return b.currentIndex
}

func (b *Board) SelectObject(idx int) error {
	if idx < 0 || idx >= len(b.objects) {
		return errors.New("index out of range")
	}

	b.selectedIndex = &idx
	return nil
}

func (b *Board) DeselectObject() {
	b.selectedIndex = nil

}
func (b *Board) GetSelectedObject() models.BoardObject {
	if b.selectedIndex == nil {
		return nil
	}

	if *b.selectedIndex < 0 || *b.selectedIndex >= len(b.objects) {
		return nil
	}

	return b.objects[*b.selectedIndex]
}
func (b *Board) GetCurrentObject() models.BoardObject {
	if b.currentIndex == nil {
		return nil
	}

	if *b.currentIndex < 0 || *b.currentIndex >= len(b.objects) {
		return nil
	}

	return b.objects[*b.currentIndex]
}

func (b *Board) GetObjects() []models.BoardObject {
	return b.objects
}

func (b *Board) IndexOf(obj models.BoardObject) int {
	if obj == nil {
		return -1
	}
	for i, o := range b.objects {
		if o == obj {
			return i
		}
	}
	return -1
}

func (b *Board) GetRound() int {
	return b.round
}
