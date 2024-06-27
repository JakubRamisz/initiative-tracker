package cmd

import (
	"fmt"
	"itr/logic"
	"itr/utils"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "itr",
	Short: "Simple initiative tracker for your tabletop RPGs",
	Run: func(cmd *cobra.Command, args []string) {
		configFlag, _ := cmd.Flags().GetString("config")
		conf, err := utils.LoadFromFile(configFlag)
		if err != nil {
			fmt.Print(err)
			return
		}
		logic.StartCombat(conf)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	dir, err := homedir.Dir()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	cfgFile := fmt.Sprintf("%s/.config/itr/config", dir)

	rootCmd.PersistentFlags().StringP("config", "c", cfgFile, "config file path")
}
