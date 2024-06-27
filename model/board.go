package model

import (
	"errors"
	"sort"
	"strings"
)

type Board struct {
	objects []BoardObject
}

func NewBoard(pcNames []string) *Board {
	board := Board{}
	for _, pc := range pcNames {
		board.AddPC(
			&Creature{
				Name: pc,
			},
		)
	}
	return &board
}

func (b *Board) AddPC(p *Creature) {
	b.objects = append(b.objects, p)
	sort.Sort(GameObjects(b.objects))
}

func (g *Board) AddNPC(n *NPC) {
	g.objects = append(g.objects, n)
	sort.Sort(GameObjects(g.objects))
}

func (b *Board) AddEvent(e *Event) {
	b.objects = append(b.objects, e)
	sort.Sort(GameObjects(b.objects))
}

func (b *Board) Remove(name string) {
	for i, obj := range b.objects {
		if strings.EqualFold(obj.GetName(), name) {
			b.objects = append(b.objects[:i], b.objects[i+1:]...)
			return
		}
	}
}

func (b *Board) GetObject(name string) *BoardObject {
	for _, obj := range b.objects {
		if strings.EqualFold(obj.GetName(), name) {
			return &obj
		}
	}
	return nil
}

func (b *Board) PCs() (pcs []*Creature) {
	for _, o := range b.objects {
		if c, ok := o.(*Creature); ok {
			pcs = append(pcs, c)
		}
	}
	return pcs
}

func (b *Board) MaxInitiative() int {
	if len(b.objects) == 0 {
		return 0
	}
	return b.objects[0].GetInitiative()
}

func (b *Board) MinInitiative() int {
	if len(b.objects) == 0 {
		return 0
	}
	return b.objects[len(b.objects)-1].GetInitiative()
}

func (b *Board) SetInitiative(name string, initiative int) {
	maxPriority := 0
	foundIdx := -1
	for i, obj := range b.objects {
		if strings.EqualFold(obj.GetName(), name) {
			foundIdx = i
			continue
		}
		if obj.GetPriority() > maxPriority {
			maxPriority = obj.GetPriority()
		}

	}

	b.objects[foundIdx].SetInitiative(initiative)
	b.objects[foundIdx].SetPriority(maxPriority + 1)
	sort.Sort(GameObjects(b.objects))
}

func (b *Board) GetObjects() []BoardObject {
	return b.objects
}

func (b *Board) IndexOf(obj BoardObject) int {
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

func (b *Board) Next(current *BoardObject) (*BoardObject, error) {
	if len(b.objects) == 0 {
		return nil, nil
	}
	if current == nil {
		return &b.objects[0], nil
	}

	for i, o := range b.objects {
		if o == *current {
			if i+1 < len(b.objects) {
				return &b.objects[i+1], nil
			}
			return &b.objects[0], nil
		}
	}
	return nil, errors.New("next object not found")
}

type GameObjects []BoardObject

func (g GameObjects) Len() int {
	return len(g)
}

func (g GameObjects) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g GameObjects) Less(i, j int) bool {
	return g[i].CompareInitiative(g[j]) > 0
}
