package views

import (
	"errors"
	"github.com/awesome-gocui/gocui"
)

const (
	O = "output"
)

func CreateOutputView(g *gocui.Gui) (*gocui.View, error) {
	tw, th := g.Size()
	outputView, err := g.SetView(O, tw-iWidth, iHeight+1, tw-1, th-hHeight-2, 0)
	// ErrUnknownView is not a real error condition.
	// It just says that the views did not exist before and needs initialization.
	if err != nil && !errors.Is(err, gocui.ErrUnknownView) {
		return nil, err
	}
	outputView.Wrap = true

	return outputView, nil

}
func UpdateOutputViewLayout(g *gocui.Gui) error {
	tw, th := g.Size()
	if _, err := g.SetView(O, tw-iWidth, iHeight+1, tw-1, th-hHeight-2, 0); err != nil {
		return err
	}

	return nil
}

func OutputMessage(g *gocui.Gui, message string) error {
	ov, err := g.View(O)
	if err != nil {
		return err
	}
	ov.Clear()
	_, err = ov.Write([]byte(message))

	return nil
}
