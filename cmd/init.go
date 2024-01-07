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
	headStatus := utils.CreateFile("test/.gitpot/HEAD")
	if headStatus == "File already exists" {
		return
	}
	config := utils.CreateFile("test/.gitpot/config")
	if config == "File already exists" {
		return
	}
	indexStatus := utils.CreateFile("test/.gitpot/index")
	if indexStatus == "File already exists" {
		return
	}
	objectsStatus := utils.CreateDir("test/.gitpot/objects")
	if objectsStatus == "Directory already exists" {
		return
	}
}

var intiCMD = &cobra.Command{
	Use:   "init",
	Short: "initializes a gitpot repository",
	Long:  "initializes a gitpot repository",
	Run: func(cmd *cobra.Command, args []string) {
		gitpotInitHandler()
	},
}
