package views

import (
	"github.com/awesome-gocui/gocui"
	"itr/logic"
)

func CursorDown(g *gocui.Gui, v *gocui.View) error {
	// Check to make sure data exists in the next line,
	// otherwise disallow scroll down.
	if v == nil || v.Name() != B {
		return nil
	}
	if logic.BoardState.GetSelectedObject() == nil {
		if err := v.SetCursor(0, 0); err != nil {
			return err
		}
	} else {
		v.MoveCursor(0, 1)
	}

	_, cy := v.Cursor()
	n, _ := v.Line(cy)

	if n != "" {
		if err := logic.BoardState.SelectObject(cy); err != nil {
			return err
		}
	} else {
		logic.BoardState.DeselectObject()
	}

	if err := RedrawBoard(g, v); err != nil {
		return err
	}
	return nil
}

func CursorUp(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() != B {
		return nil
	}
	if logic.BoardState.GetSelectedObject() == nil {
		if err := v.SetCursor(0, len(logic.BoardState.GetObjects())-1); err != nil {
			return err
		}
	} else {
		if _, cy := v.Cursor(); cy == 0 {
			if err := v.SetCursor(0, len(logic.BoardState.GetObjects())); err != nil {
				return err
			}
		} else {
			v.MoveCursor(0, -1)
		}
	}
	_, cy := v.Cursor()
	n, _ := v.Line(cy)
	if n != "" {
		if err := logic.BoardState.SelectObject(cy); err != nil {
			return err
		}
	} else {
		logic.BoardState.DeselectObject()
	}
	if err := RedrawBoard(g, v); err != nil {
		return err
	}
	return nil
}
