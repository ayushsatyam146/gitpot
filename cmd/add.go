package cmd

import (
	file "github.com/ayushsatyam146/gitpot/files"
	"github.com/ayushsatyam146/gitpot/index"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCMD)
}

func addHandler(args []string) {
	tree := index.BuildTree(args)
	_, indexFileContent := file.WriteTreeToGitpot(tree, "test/.gitpot")
	index.WriteContentToIndex("test/.gitpot", indexFileContent)
}

var addCMD = &cobra.Command{
	Use:   "add",
	Short: "initializes a gitpot repository",
	Long:  "initializes a gitpot repository",
	Run: func(cmd *cobra.Command, args []string) {
		addHandler(args)
	},
}
