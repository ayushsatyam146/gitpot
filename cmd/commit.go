package cmd

import (
	"fmt"

	file "github.com/ayushsatyam146/gitpot/files"
	"github.com/ayushsatyam146/gitpot/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(commitCMD)
}

func commitHandler() {
	content := file.ReadFile("test/.gitpot/index")
	contentString := "tree\n" + string(content)
	treeHash,_ := utils.GetSHA1([]byte(contentString))
	author := "Ayush Satyam"
	committer := "Ayush Satyam"
	message := "Initial Commit"
	// read parent from HEAD eventually
	parent := ""
	commitString := "commit\n" + 
	"tree " + treeHash + "\n" + 
	"author " + author + "\n" + 
	"parent" + parent + "\n" + 
	"committer " + committer + "\n\n" + 
	message + "\n"
	fmt.Println(commitString)
	commitHash := utils.WriteToObjectsDir("test/.gitpot", []byte(commitString), true)
	// update HEAD and refs/heads/branch_name along with this
	// so that relevant branch has it's latest commit hash
	// and HEAD will also point to that relevant branch's latest commit hash
	// indicating active branch
	fmt.Println(commitHash)
}

var commitCMD = &cobra.Command{
	Use:   "commit",
	Short: "initializes a gitpot repository",
	Long:  "initializes a gitpot repository",
	Run: func(cmd *cobra.Command, args []string) {
		commitHandler()
	},
}
