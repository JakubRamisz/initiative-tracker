package views

import (
	"errors"
	"github.com/awesome-gocui/gocui"
)

const (
	I       = "input"
	iWidth  = 40
	iHeight = 2
)

func CreateInputView(g *gocui.Gui) (*gocui.View, error) {
	tw, _ := g.Size()
	inputView, err := g.SetView(I, tw-iWidth, 0, tw-1, iHeight, 0)

	// ErrUnknownView is not a real error condition.
	// It just says that the views did not exist before and needs initialization.
	if err != nil && !errors.Is(err, gocui.ErrUnknownView) {
		return nil, err
	}

	inputView.Title = "Input"
	inputView.Editable = true

	return inputView, nil

}
func UpdateInputViewLayout(g *gocui.Gui) error {
	tw, _ := g.Size()
	if _, err := g.SetView(I, tw-iWidth, 0, tw-1, iHeight, 0); err != nil {
		return err
	}

	return nil
}

func Input(g *gocui.Gui, cv *gocui.View) error {
	if cv.Name() != B {
		return nil
	}

	nv, err := g.SetCurrentView(I)
	if err != nil {
		return err
	}
	g.Cursor = true
	if err = nv.SetCursor(0, 0); err != nil {
		return err
	}
	return nil
}

func GoBack(g *gocui.Gui, cv *gocui.View) error {
	if cv.Name() != I {
		return nil
	}
	cv.Clear()
	_, err := g.SetCurrentView(B)
	if err != nil {
		return err
	}
	g.Cursor = false
	return nil
}
