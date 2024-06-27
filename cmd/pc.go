package cmd

import (
	"fmt"
	"itr/utils"
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
		conf, err := utils.LoadFromFile(configFlag)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		pcString := strings.Join(conf.PlayerCharacters, "\n")
		fmt.Println(pcString)
	},
}
