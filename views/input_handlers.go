package views

import (
	"errors"
	"fmt"
	"github.com/awesome-gocui/gocui"
	"itr/logic"
	"itr/models"
	"strconv"
	"strings"
)

func HandleInput(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	line, _ := v.Line(cy)

	v.Clear()
	if err := v.SetCursor(0, 0); err != nil {
		return err
	}

	cmd := strings.Split(strings.TrimSpace(line), " ")

	str, err := handleCommand(cmd)
	if err != nil {
		if err := OutputMessage(g, "Error: "+err.Error()); err != nil {
			return err
		}
		return nil
	} else {
		if err := OutputMessage(g, str); err != nil {
			return err
		}
	}
	bv, err := g.View(B)
	if err != nil {
		return err
	}
	if err := RedrawBoard(g, bv); err != nil {
		return err
	}

	return GoBack(g, v)
}

func handleCommand(cmd []string) (string, error) {
	switch cmd[0] {
	case "p":
		if len(cmd) < 2 {
			err := errors.New("provide a name for the player character")
			return "", err
		}

		name := strings.Join(cmd[1:], " ")
		if err := logic.BoardState.AddPC(name); err != nil {
			return "", err
		}
		return fmt.Sprintf("Added PC: %s", name), nil

	case "n":
		if len(cmd) < 3 {
			err := errors.New("provide a name and hp for the npc")
			return "", err
		}

		name := strings.Join(cmd[1:len(cmd)-1], " ")
		hp, err := strconv.Atoi(cmd[len(cmd)-1])
		if err != nil {
			return "", err
		}
		if err := logic.BoardState.AddNPC(name, hp); err != nil {
			return "", err
		}
		return fmt.Sprintf("Added NPC: %s", name), nil

	case "e":
		if len(cmd) < 3 {
			err := errors.New("provide a name and initiative for the event")
			return "", err
		}

		name := strings.Join(cmd[1:len(cmd)-1], " ")
		in, err := strconv.Atoi(cmd[len(cmd)-1])
		if err != nil {
			return "", err
		}
		if err := logic.BoardState.AddEvent(name, in); err != nil {
			return "", err
		}
		return fmt.Sprintf("Added event: %s", name), nil

	case "r":
		obj := logic.BoardState.GetSelectedObject()
		if obj == nil {
			return "", errors.New("no object selected")
		}
		if err := logic.BoardState.RemoveObject(obj); err != nil {
			return "", err
		}
		return fmt.Sprintf("Removed: %s", obj.GetName()), nil

	case "c":
		obj := logic.BoardState.GetSelectedObject()
		if obj == nil {
			return "", errors.New("no object selected")
		}
		if err := logic.BoardState.CloneObject(obj); err != nil {
			return "", err
		}
		return fmt.Sprintf("Cloned %s", obj.GetName()), nil

	case "i":
		if len(cmd) < 2 {
			err := errors.New("provide an initiative value")
			return "", err
		}
		obj := logic.BoardState.GetSelectedObject()
		if obj == nil {
			return "", errors.New("no object selected")
		}
		in, err := strconv.Atoi(cmd[len(cmd)-1])
		if err != nil {
			return "", err
		}
		if err := logic.BoardState.SetInitiative(obj, in); err != nil {
			return "", err
		}
		return fmt.Sprintf("Set initiative for %s to %d", obj.GetName(), in), nil

	case "d":
		if len(cmd) != 2 {
			err := errors.New("provide the HP amount")
			return "", err
		}
		amount, err := strconv.Atoi(cmd[1])
		if err != nil {
			err := errors.New("provide the HP amount")
			return "", err
		}

		obj := logic.BoardState.GetSelectedObject()
		if obj == nil {
			return "", errors.New("no object selected")
		}
		npc, ok := obj.(*models.NPC)
		if !ok {
			return "", errors.New("only NPCs can take damage")
		}
		npc.RecieveDMG(amount)
		return fmt.Sprintf("Dealt %d damage to %s", amount, obj.GetName()), nil

	case "h":
		if len(cmd) != 2 {
			err := errors.New("provide the HP amount")
			return "", err
		}
		amount, err := strconv.Atoi(cmd[1])
		if err != nil {
			err := errors.New("provide the HP amount")
			return "", err
		}

		obj := logic.BoardState.GetSelectedObject()
		if obj == nil {
			return "", errors.New("no object selected")
		}
		npc, ok := obj.(*models.NPC)
		if !ok {
			return "", errors.New("only NPCs can heal")
		}
		npc.Heal(amount)
		return fmt.Sprintf("Healed %d to %s", amount, obj.GetName()), nil

	case "s":
		err := logic.BoardState.Step()
		if err != nil {
			return "", err
		}
		return "Moved to next initiative count", nil

	default:
		return "", errors.New("unknown command: " + cmd[0])
	}
}
