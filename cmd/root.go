package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gitpot",
		Short: "gitpot is a simple implementation of git in golang",
		Long:  "gitpot is a simple implementation of git in golang",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
