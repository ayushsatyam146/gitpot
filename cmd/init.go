package cmd

import (
	// utils "github.com/ayushsatyam146/gitpot/utils"
	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(intiCMD)
}

func gitpotInitHandler() {
	// init .gitpot directory from here
}

var intiCMD = &cobra.Command{
	Use:   "init",
	Short: "initializes a gitpot repository",
	Long:  "initializes a gitpot repository",
	Run: func(cmd *cobra.Command, args []string) {
		gitpotInitHandler()
	},
}