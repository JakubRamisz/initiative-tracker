package views

import (
	"github.com/awesome-gocui/gocui"
)

func Keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, Quit); err != nil {
		return err
	}
	if err := g.SetKeybinding(B, gocui.KeyArrowUp, gocui.ModNone, CursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding(B, gocui.KeyArrowDown, gocui.ModNone, CursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding(B, gocui.KeySpace, gocui.ModNone, Input); err != nil {
		return err
	}
	if err := g.SetKeybinding(I, gocui.KeyEsc, gocui.ModNone, GoBack); err != nil {
		return err
	}
	if err := g.SetKeybinding(I, gocui.KeyEnter, gocui.ModNone, HandleInput); err != nil {
		return err
	}

	return nil
}
