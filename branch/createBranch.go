package branch

import (
	"fmt"
	"os"

	file "github.com/ayushsatyam146/gitpot/files"
)

func CreateBranch(branchName string) {
	dirPath := "test/.gitpot/refs/heads"
	currentBranch := string(file.ReadFile("test/.gitpot/HEAD"))
	currentLatestHash := string(file.ReadFile("test/.gitpot/" + currentBranch))
	branchFile := dirPath + "/" + branchName
	file, err := os.Create(branchFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	file.WriteString(currentLatestHash)
}