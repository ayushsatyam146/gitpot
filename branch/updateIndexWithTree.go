package branch

import (
	"fmt"
	"os"
	"strings"

	"github.com/ayushsatyam146/gitpot/files"
)

func UpdateIndexWithTree(treeHash string) {
	treeContent := string(files.ReadFile("test/.gitpot/objects/" + treeHash[:2] + "/" + treeHash[2:]))
	indexContent := ""
	for index , line := range strings.Split(treeContent, "\n") {
		if index != 0 {
			indexContent += line + "\n"
		}
	}
	file, err := os.Open("test/.gitpot/index")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	file.WriteString(indexContent)
}