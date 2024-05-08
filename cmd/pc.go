package cmd

import (
	"fmt"
	"itr/logic"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pcCmd)
}

var pcCmd = &cobra.Command{
	Use:   "pc",
	Short: "List added player characters",
	Run: func(cmd *cobra.Command, args []string) {
		configFlag, _ := cmd.Flags().GetString("config")
		config, err := logic.ReadConfig(configFlag)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		pcString := strings.Join(config.PlayerCharacters, "\n")
		fmt.Println(pcString)
	},
}
