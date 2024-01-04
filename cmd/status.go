package cmd

import (
	// file "github.com/ayushsatyam146/gitpot/files"
	// "github.com/ayushsatyam146/gitpot/index"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCMD)
}

func statusHandler(args []string) {
	fmt.Println("status")
}

var statusCMD = &cobra.Command{
	Use:   "status",
	Short: "gives you the status of tracked/untracked files as well as list of changes in the repository",
	Long:  "gives you the status of tracked/untracked files as well as list of changes in the repository",
	Run: func(cmd *cobra.Command, args []string) {
		statusHandler(args)
	},
}
