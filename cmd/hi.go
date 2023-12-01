package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(hiCmd)
}

func hiHandler() {
	fmt.Println("say hello")
}

var hiCmd = &cobra.Command{
	Use:   "ls",
	Short: "listing command",
	Long:  "listing command",
	Run: func(cmd *cobra.Command, args []string) {
		hiHandler()
	},
}