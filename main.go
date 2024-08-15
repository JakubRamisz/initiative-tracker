package main

import (
	"fmt"
	"itr/logic"
	"itr/views"
	"log"
	"os"
	"os/user"

	"github.com/awesome-gocui/gocui"
)

func main() {
	usr, _ := user.Current()
	dir := usr.HomeDir
	configPath := dir + "/.config/itr/config.json"
	logic.ConfigFilePath = configPath

	// Log file
	f, err := os.OpenFile(dir+"/.itr.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer f.Close()
	log.SetOutput(f)

	g, err := gocui.NewGui(gocui.Output256, true)
	if err != nil {
		log.Println("Failed to create a GUI:", err)
		return
	}
	defer g.Close()

	g.Highlight = true
	g.SelFrameColor = gocui.ColorCyan

	g.SetManagerFunc(layout)

	// Create views
	boardView, err := views.CreateBoardView(g)
	if err != nil {
		log.Println("Failed to create board view:", err)
		return
	}
	if _, err = views.CreateInputView(g); err != nil {
		log.Println("Failed to create input view:", err)
		return
	}
	if _, err = views.CreateOutputView(g); err != nil {
		log.Println("Failed to create output view:", err)
		return
	}
	if _, err = views.CreateHelpView(g); err != nil {
		log.Println("Failed to create help view:", err)
		return
	}

	// Load config
	conf, err := logic.LoadFromFile(configPath)
	if err != nil {
		log.Println("Failed to load config:", err)
		return
	}
	logic.BoardState = logic.NewBoard(conf.PlayerCharacters)
	if len(logic.BoardState.PCs()) > 0 {
		if err := views.RedrawBoard(g, boardView); err != nil {
			log.Println("Failed to redraw board:", err)
			return
		}
	}

	// Apply keybindings to program.
	if err = views.Keybindings(g); err != nil {
		log.Println("Failed to set keybindings: ", err)
		return
	}

	// Set initial views here, right before program start
	if _, err := g.SetCurrentView(views.B); err != nil {
		log.Println("Failed to set current view to board:", err)
		return
	}

	// Start the main loop.
	err = g.MainLoop()
	log.Println("Main loop has finished:", err)
}

func layout(g *gocui.Gui) error {
	// Update the views according to the new terminal size.
	if err := views.UpdateBoardViewLayout(g); err != nil {
		return err
	}
	if err := views.UpdateInputViewLayout(g); err != nil {
		return err
	}
	if err := views.UpdateOutputViewLayout(g); err != nil {
		return err
	}
	if err := views.UpdateHelpViewLayout(g); err != nil {
		return err
	}

	return nil
}
