package logic

import (
	"fmt"
	"itr/model"
	"strconv"

	tm "github.com/buger/goterm"
	"github.com/fatih/color"
)

func (c *Combat) displayInitiativeOrder() {
	tm.Clear()
	initiativeStr := "-"
	if c.currentObject != nil {
		obj := *c.currentObject
		initiativeStr = fmt.Sprintf("%d", obj.GetInitiative())
	}
	objects := c.board.GetObjects()
	tm.Printf("----------- Round: %d Initiative: %s -----------\n", c.round, initiativeStr)

	for _, o := range objects {
		str := fmt.Sprintf("%d\t%s", o.GetInitiative(), o.GetInfo())
		if npc, ok := o.(*model.NPC); ok && npc.CurrentHP == 0 {
			str = color.HiBlackString(str)

		} else if c.currentObject != nil && o == *c.currentObject {
			str = tm.Background(str, tm.BLUE)
		}
		tm.Println(str)
	}
	tm.Println()
	tm.Println(color.HiBlackString(helpMessage))
	tm.Flush()
}

func (c *Combat) findNextObject(currentObject *model.BoardObject) (*model.BoardObject, error) {
	object, err := c.board.Next(currentObject)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if object == nil {
		return nil, nil
	}
	o := *object
	if obj, ok := o.(*model.NPC); ok && obj.CurrentHP <= 0 {
		return c.findNextObject(object)
	}
	return object, nil
}

func (c *Combat) getAdjustHPParams(params []string) (int, *model.NPC) {
	if len(params) != 2 {
		fmt.Println("Invalid arguments. Please provide the name of the creature and the HP amount")
		return 0, nil
	}

	name := params[0]
	amount, err := strconv.Atoi(params[1])
	if err != nil {
		fmt.Println("Invalid arguments. Please provide the name of the creature and the HP amount")
		return 0, nil
	}

	npc := c.board.GetObject(name)
	if npc == nil {
		fmt.Println("NPC not found")
		return 0, nil
	}
	obj, ok := (*npc).(*model.NPC)
	if !ok {
		fmt.Println("NPC not found")
		return 0, nil
	}
	return amount, obj
}
