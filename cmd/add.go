package cmd

import (
	file "github.com/ayushsatyam146/gitpot/files"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCMD)
}

func addHandler(args []string) {
	// file.PrintTree(file.GetRelTreeFromWorkingDir(args[0]))
	file.GetRelTreeFromWorkingDir(args[0])
}

var addCMD = &cobra.Command{
	Use:   "add",
	Short: "initializes a gitpot repository",
	Long:  "initializes a gitpot repository",
	Run: func(cmd *cobra.Command, args []string) {
		addHandler(args)
	},
}
