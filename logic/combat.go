package logic

import (
	"bufio"
	"fmt"
	"itr/model"
	"itr/utils"
	"os"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

type Combat struct {
	board         *model.Board
	currentObject *model.BoardObject
	round         int
}

func StartCombat(conf utils.Config) {
	pcNames := conf.PlayerCharacters
	cmd := model.Command{}
	board := model.NewBoard(pcNames)
	c := Combat{
		board: board,
		round: 1,
	}

	for {
		switch cmd.Code {
		case cmdCodeWelcome:
			cmd = c.cmdWelcome()
		case cmdCodeHelp:
			cmd = c.cmdHelp()
		case cmdCodeAddNPC:
			cmd = c.cmdAddNPC(cmd.Params)
		case cmdCodeAddEvent:
			cmd = c.cmdAddEvent(cmd.Params)
		case cmdCodeRemove:
			cmd = c.cmdRemove(cmd.Params)
		case cmdSetInitiative:
			cmd = c.cmdSetInitiative(cmd.Params)
		case cmdCodeNextStep:
			cmd = c.cmdNextStep()
		case cmdCodeDamage:
			cmd = c.cmdDamage(cmd.Params)
		case cmdCodeHeal:
			cmd = c.cmdHeal(cmd.Params)
		case cmdCodeExit:
			syscall.Exit(0)
		}
	}
}

func getCommand() model.Command {
	fmt.Print("> ")

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		color.HiRed("Error reading input: %s\n", err)
		return getCommand()
	}
	input := strings.Split(strings.TrimSpace(scanner.Text()), " ")

	if len(input) == 0 || input[0] == "" {
		return getCommand()
	}

	c := model.Command{}
	if len(input) > 1 {
		c.Params = input[1:]
	}
	input[0] = strings.ToLower(input[0])

	switch input[0] {
	case "help":
		c.Code = cmdCodeHelp
	case "exit":
		c.Code = cmdCodeExit
	case "npc":
		c.Code = cmdCodeAddNPC
	case "ev":
		c.Code = cmdCodeAddEvent
	case "rm":
		c.Code = cmdCodeRemove
	case "i":
		c.Code = cmdSetInitiative
	case "s":
		c.Code = cmdCodeNextStep
	case "dmg":
		c.Code = cmdCodeDamage
	case "heal":
		c.Code = cmdCodeHeal
	default:
		color.HiRed("Invalid command: %s\n", input[0])
		return getCommand()

	}
	return c
}
