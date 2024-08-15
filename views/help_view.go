package views

import (
	"errors"
	"fmt"
	"github.com/awesome-gocui/gocui"
)

const (
	H       = "help"
	hHeight = 9
)

func CreateHelpView(g *gocui.Gui) (*gocui.View, error) {
	tw, th := g.Size()
	helpView, err := g.SetView(H, tw-iWidth, th-hHeight-1, tw-1, th-1, 0)
	// ErrUnknownView is not a real error condition.
	// It just says that the views did not exist before and needs initialization.
	if err != nil && !errors.Is(err, gocui.ErrUnknownView) {
		return nil, err
	}

	helpView.Title = "Commands"

	commands := []string{
		"p <name> - add a pc",
		"n <name> <hp> - add an npc",
		"e <name> <initiative> - add an event",
		"r - remove object",
		"i <initiative> - set initiative",
		"s - next initiative count",
		"d <amount> - deal damage",
		"h <amount> - heal",
	}

	for _, command := range commands {
		if _, err = fmt.Fprintln(helpView, command); err != nil {
			return nil, err
		}
	}
	return helpView, nil

}
func UpdateHelpViewLayout(g *gocui.Gui) error {
	tw, th := g.Size()
	if _, err := g.SetView(H, tw-iWidth, th-hHeight-1, tw-1, th-1, 0); err != nil {
		return err
	}

	return nil
}
