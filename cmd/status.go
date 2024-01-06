package cmd

import (
	"github.com/ayushsatyam146/gitpot/status"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCMD)
}

func statusHandler(args []string) {
	status.PrintUntrackedFiles()
	status.PrintModifiedFiles()
}

var statusCMD = &cobra.Command{
	Use:   "status",
	Short: "gives you the status of tracked/untracked files as well as list of changes in the repository",
	Long:  "gives you the status of tracked/untracked files as well as list of changes in the repository",
	Run: func(cmd *cobra.Command, args []string) {
		statusHandler(args)
	},
}

