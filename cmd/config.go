package cmd

import (
	"github.com/ayushsatyam146/gitpot/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(configCMD)
}

func configHandler(args []string) {
	if len(args) == 0 {
		utils.ReadFromConfig("user.name")
		utils.ReadFromConfig("user.email")
		return
	}
	if len(args) != 2 {
		panic("Invalid number of arguments")
	}
	if args[0] != "user.name" && args[0] != "user.email" {
		panic("Invalid config key")
	}
	utils.WriteToConfig(args[0], args[1])
}

var configCMD = &cobra.Command{
	Use:   "config",
	Short: "takes the latest updates to the index or staging area and commits them to the repository",
	Long:  "takes the latest updates to the index or staging area and commits them to the repository",
	Run: func(cmd *cobra.Command, args []string) {
		configHandler(args)
	},
}
