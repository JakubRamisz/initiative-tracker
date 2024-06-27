package cmd

import (
	"errors"
	"fmt"
	"itr/utils"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pcAddCmd)
}

var pcAddCmd = &cobra.Command{
	Use:   "pc-add",
	Short: "Adds a player character",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one player character name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		configFlag, _ := cmd.Flags().GetString("config")
		err := addPCs(args, configFlag)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	},
}

func addPCs(pcs []string, configPath string) error {
	conf, err := utils.LoadFromFile(configPath)
	if err != nil {
		return err
	}

	for _, pc := range pcs {
		pc := pc
		if slices.Contains(conf.PlayerCharacters, pc) {
			continue
		}

		conf.PlayerCharacters = append(conf.PlayerCharacters, pc)
	}

	err = conf.SaveToFile(configPath)
	return err
}
