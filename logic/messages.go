package logic

const (
	menuHelpMessage = `Commands:
	- help: displays this message
	- npc <name> <hitpoints>: adds an npc
	- ev <name> <initiative>: adds an event (defaults: name: Lair Action  initiative: 20)
	- rm <name>: removes an object
	- i <name> <initiative>: sets the initiative of a creature
	- s: advances to the next initiative count
	- dmg: <name> <amount>: deals damage to a creature
	- heal: <name> <amount>: heals a creature
	- exit: exits the program
	`

	helpMessage = "Type 'help' for a list of commands."
)
