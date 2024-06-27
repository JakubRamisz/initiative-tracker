package logic

import (
	"fmt"
	"itr/model"
	"strconv"
	"strings"
)

const (
	cmdCodeWelcome = iota
	cmdCodeHelp
	cmdCodeAddNPC
	cmdCodeAddEvent
	cmdCodeRemove
	cmdCodeNextRound
	cmdSetInitiative
	cmdCodeNextStep
	cmdCodeDamage
	cmdCodeHeal
	cmdCodeExit
)

func (c *Combat) cmdHelp() model.Command {
	fmt.Println(menuHelpMessage)
	return getCommand()
}

func (c *Combat) cmdWelcome() model.Command {
	c.displayInitiativeOrder()
	return getCommand()
}

func (c *Combat) cmdAddNPC(params []string) model.Command {
	if len(params) < 2 {
		fmt.Println("Invalid arguments. Please provide a name and HP for the creature")
		return getCommand()
	}

	name := strings.Join(params[:len(params)-1], " ")
	hp, err := strconv.Atoi(params[len(params)-1])
	if err != nil {
		fmt.Println("Invalid arguments. Please provide a name and HP for the creature")
		return getCommand()
	}

	objects := c.board.GetObjects()
	exists := 0
	for _, obj := range objects {
		if obj.GetName() == name {
			exists++
		}
	}
	if exists > 0 {
		name = fmt.Sprintf("%s (%d)", name, exists)
	}

	npc := model.NPC{
		Creature: model.Creature{
			Name: name,
		},
		MaxHP:     hp,
		CurrentHP: hp,
	}
	c.board.AddNPC(&npc)

	c.displayInitiativeOrder()
	return getCommand()
}

func (c *Combat) cmdRemove(params []string) model.Command {
	if len(params) < 1 {
		fmt.Println("Invalid arguments. Please provide the name of the creature to remove")
		return getCommand()
	}

	name := strings.Join(params, " ")
	c.board.Remove(name)

	c.displayInitiativeOrder()
	return getCommand()
}

func (c *Combat) cmdSetInitiative(params []string) model.Command {
	if len(params) < 2 {
		fmt.Println("Invalid arguments. Please provide the name of the creature and its initiative")
		return getCommand()
	}

	name := strings.Join(params[:len(params)-1], " ")
	initiative, err := strconv.Atoi(params[len(params)-1])
	if err != nil {
		fmt.Println("Invalid arguments. Please provide the name of the creature and its initiative")
		return getCommand()
	}

	c.board.SetInitiative(name, initiative)

	c.displayInitiativeOrder()
	return getCommand()
}

func (c *Combat) cmdNextStep() model.Command {
	nextRound := false
	if c.currentObject != nil && c.board.IndexOf(*c.currentObject) == len(c.board.GetObjects())-1 {
		nextRound = true
	}

	object, err := c.findNextObject(c.currentObject)
	if err != nil {
		fmt.Println(err)
		return getCommand()
	}

	c.currentObject = object
	if nextRound {
		c.round++
	}

	c.displayInitiativeOrder()
	return getCommand()
}

func (c *Combat) cmdDamage(params []string) model.Command {
	amount, obj := c.getAdjustHPParams(params)
	if obj == nil {
		return getCommand()
	}
	obj.RecieveDMG(amount)

	c.displayInitiativeOrder()
	return getCommand()
}

func (c *Combat) cmdHeal(params []string) model.Command {
	amount, obj := c.getAdjustHPParams(params)
	if obj == nil {
		return getCommand()
	}
	obj.Heal(amount)

	c.displayInitiativeOrder()
	return getCommand()
}

func (c *Combat) cmdAddEvent(params []string) model.Command {
	name := "Lair Action"
	initiative := 20
	if len(params) >= 2 {
		name = strings.Join(params[:len(params)-1], " ")
		in, err := strconv.Atoi(params[len(params)-1])
		if err != nil {
			fmt.Println("Invalid arguments. Please provide the name of the event and its initiative")
			return getCommand()
		}
		initiative = in
	}

	event := model.Event{
		Name:       name,
		Initiative: initiative,
	}
	c.board.AddEvent(&event)

	c.displayInitiativeOrder()
	return getCommand()
}
