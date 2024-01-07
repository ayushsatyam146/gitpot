package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(branchCMD)
}

func branchHandler(args []string) {
	if len(args) == 0 {
		// PrintBrancheNames()
	} else {
		// CreateBranch(args[0])
	}
	// no args then list all branches
	// if args then create a new branch
}

var branchCMD = &cobra.Command{
	Use:   "add",
	Short: "adds the listed files or directories to the staging area",
	Long:  "adds the listed files or directories to the staging area",
	Run: func(cmd *cobra.Command, args []string) {
		branchHandler(args)
	},
}
