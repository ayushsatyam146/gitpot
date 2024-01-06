package cmd

import (
	"os"

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
	treeHash, _ := utils.GetSHA1([]byte(contentString))
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
	utils.WriteToObjectsDir("test/.gitpot", []byte(commitString), true)
	commitHash,_ := utils.GetSHA1([]byte(commitString))
	file, err := os.Create("test/.gitpot/HEAD")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(commitHash)
	// update HEAD and refs/heads/branch_name along with this
	// so that relevant branch has it's latest commit hash
	// and HEAD will also point to that relevant branch's latest commit hash
	// indicating active branch
}

var commitCMD = &cobra.Command{
	Use:   "commit",
	Short: "takes the latest updates to the index or staging area and commits them to the repository",
	Long:  "takes the latest updates to the index or staging area and commits them to the repository",
	Run: func(cmd *cobra.Command, args []string) {
		commitHandler()
	},
}
