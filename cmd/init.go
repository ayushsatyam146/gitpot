package cmd

import (
	"fmt"

	"github.com/ayushsatyam146/gitpot/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(intiCMD)
}

func gitpotInitHandler() {
	status := utils.CreateDir("test/.gitpot")
	if status == "Directory already exists" {
		fmt.Println("Gitpot repository already initialized")
		return 
	} else {
		fmt.Println("Initialized empty Gitpot repository in test/.gitpot")
	}
	utils.CreateFile("test/.gitpot/HEAD", []byte("refs/heads/master"))	
	utils.CreateFile("test/.gitpot/config", []byte("[core]\n\trepositoryformatversion = 0\n\tfilemode = true\n\tbare = false\n\tlogallrefupdates = true\n"))
	utils.CreateFile("test/.gitpot/index", []byte(""))
	utils.CreateDir("test/.gitpot/objects")
	utils.CreateFile("test/.gitpot/refs/heads/master", []byte(""))
}

var intiCMD = &cobra.Command{
	Use:   "init",
	Short: "initializes a gitpot repository",
	Long:  "initializes a gitpot repository",
	Run: func(cmd *cobra.Command, args []string) {
		gitpotInitHandler()
	},
}
