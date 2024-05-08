package cmd

import (
	"errors"
	"fmt"
	"itr/logic"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pcRemoveCmd)
}

var pcRemoveCmd = &cobra.Command{
	Use:   "pc-remove",
	Short: "Removes a player character",

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one player character name")
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		configFlag, _ := cmd.Flags().GetString("config")
		err := removePCs(args, configFlag)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	},
}

func removePCs(pcs []string, configPath string) error {
	config, err := logic.ReadConfig(configPath)
	if err != nil {
		return err
	}
	newPCs := []string{}
	for _, pc := range config.PlayerCharacters {
		pc := pc
		if slices.Contains(pcs, pc) {
			continue
		}

		newPCs = append(newPCs, pc)
	}
	config.PlayerCharacters = newPCs

	err = logic.WriteConfig(config, configPath)
	return err
}
