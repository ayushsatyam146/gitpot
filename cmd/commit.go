package cmd

import (
	"time"

	"github.com/ayushsatyam146/gitpot/branch"
	file "github.com/ayushsatyam146/gitpot/files"
	"github.com/ayushsatyam146/gitpot/utils"
	"github.com/spf13/cobra"
)

var commitMessage string

func init() {
	rootCmd.AddCommand(commitCMD)
	commitCMD.Flags().StringVar(&commitMessage, "m", "", "write a commit message")
}

func commitHandler() {
	content := file.ReadFile("test/.gitpot/index")
	contentString := "tree\n" + string(content)
	treeHash, _ := utils.GetSHA1([]byte(contentString))
	committer := utils.ReadFromConfig("user.name") + " <" + utils.ReadFromConfig("user.email") + ">"
	parentCommit := branch.GetLatestCommit()
	Date := time.Now().Format("Mon Jan 2 15:04:05 2006 -0700")
	commitString := "commit\n" +
		"tree " + treeHash + "\n" +
		"parent " + parentCommit + "\n" +
		"date " + Date + "\n" +
		"committer " + committer + "\n\n" +
		commitMessage + "\n"
	utils.WriteToObjectsDir("test/.gitpot", []byte(commitString), true)
	commitHash,_ := utils.GetSHA1([]byte(commitString))
	utils.UpdateCommitHashInBranch(commitHash)
}

var commitCMD = &cobra.Command{
	Use:   "commit",
	Short: "takes the latest updates to the index or staging area and commits them to the repository",
	Long:  "takes the latest updates to the index or staging area and commits them to the repository",
	Run: func(cmd *cobra.Command, args []string) {
		commitHandler()
	},
}
