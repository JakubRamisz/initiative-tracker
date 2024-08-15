package views

import (
	"errors"
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/fatih/color"
	"itr/logic"
	"itr/models"
)

const (
	B = "board"
)

func CreateBoardView(g *gocui.Gui) (*gocui.View, error) {
	tw, th := g.Size()
	boardView, err := g.SetView(B, 0, 0, tw-iWidth-1, th-1, 0)

	// ErrUnknownView is not a real error condition.
	// It just says that the views did not exist before and needs initialization.
	if err != nil && !errors.Is(err, gocui.ErrUnknownView) {
		return nil, err
	}
	boardView.Title = ""

	return boardView, nil
}

func UpdateBoardViewLayout(g *gocui.Gui) error {
	tw, th := g.Size()
	v, err := g.SetView(B, 0, 0, tw-iWidth-1, th-1, 0)
	if err != nil {
		return err
	}

	initiativeStr := ""
	obj := logic.BoardState.GetCurrentObject()
	if obj != nil {
		s := obj.GetInitiative()
		initiativeStr = fmt.Sprintf("%d", s)
	}
	v.Title = fmt.Sprintf("Round: %d Initiative: %s ", logic.BoardState.GetRound(), initiativeStr)

	return nil
}

func RedrawBoard(g *gocui.Gui, v *gocui.View) error {
	v.Clear()
	items := logic.BoardState.GetObjects()

	for i, item := range items {
		str := item.GetInfo()
		currentObjectStr := ""
		c := color.New(color.FgWhite)

		// Selected item color
		selectedIndex := logic.BoardState.GetSelectedIndex()
		if selectedIndex != nil && *selectedIndex == i {
			c = c.Add(color.BgWhite, color.FgBlack)
		}

		// Current item color
		currentIndex := logic.BoardState.GetCurrentIndex()
		if currentIndex != nil && *currentIndex == i {
			currentObjectStr = ">"
			c = c.Add(color.Bold, color.Underline)
		}

		if npc, ok := item.(*models.NPC); ok && npc.CurrentHP == 0 {
			c = c.Add(color.CrossedOut)
		}

		_, err := c.Fprintln(v, fmt.Sprintf("%1s %-300s", currentObjectStr, str))
		if err != nil {
			return err
		}
	}
	if idx := logic.BoardState.GetSelectedIndex(); idx != nil {
		if err := v.SetCursor(0, *idx); err != nil {
			return err
		}
	}
	return nil
}
