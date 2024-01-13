package branch

import (
	"fmt"
	"os"
)

func SwitchHead(branchName string) {
	file, err := os.Create("test/.gitpot/HEAD")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()
	file.WriteString("refs/heads/" + branchName)
}

