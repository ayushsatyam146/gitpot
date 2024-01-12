package cmd

import (
	"github.com/ayushsatyam146/gitpot/branch"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(branchCMD)
}

func branchHandler(args []string) {
	if len(args) == 0 {
		branch.PrintBrancheNames()
	} else {
		branch.CreateBranch(args[0])
	}
}

var branchCMD = &cobra.Command{
	Use:   "branch",
	Short: "adds the listed files or directories to the staging area",
	Long:  "adds the listed files or directories to the staging area",
	Run: func(cmd *cobra.Command, args []string) {
		branchHandler(args)
	},
}
