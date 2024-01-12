package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ayushsatyam146/gitpot/files"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkoutCMD)
}

func checkoutHandler(args []string) {
	branchName := args[0]
	branchFile := "test/.gitpot/refs/heads/" + branchName
	if _, err := os.Stat(branchFile); os.IsNotExist(err) {
		fmt.Println("Branch doesn't exist")
	} else {
		file, err := os.Open("test/.gitpot/HEAD")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		file.WriteString("refs/heads" + branchName)

		commitHash := string(files.ReadFile("test/.gitpot/refs/heads/" + branchName))
		commitContent := string(files.ReadFile("test/.gitpot/objects/" + commitHash[:2] + "/" + commitHash[2:]))
		treeHash := ""
		for _, line := range strings.Split(commitContent, "\n") {
			if strings.HasPrefix(line, "tree") {
				treeHash = strings.Split(line, " ")[1]
				break
			}
		}
		tree := files.GetTreeFromHash("test/.gitpot",treeHash,"root")
		files.ClearWorkingDir()
		files.WriteTreeToWorkingDir(tree,"test")
	}

}

var checkoutCMD = &cobra.Command{
	Use:   "checkout",
	Short: "adds the listed files or directories to the staging area",
	Long:  "adds the listed files or directories to the staging area",
	Run: func(cmd *cobra.Command, args []string) {
		checkoutHandler(args)
	},
}
