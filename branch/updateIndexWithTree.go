package branch

import (
	"fmt"
	"os"
	"strings"

	"github.com/ayushsatyam146/gitpot/files"
)

func UpdateIndexWithTree(treeHash string) {
	treeContent := string(files.ReadFile("test/.gitpot/objects/" + treeHash[:2] + "/" + treeHash[2:]))
	stringContent := string(treeContent)
	element := strings.Split(stringContent, "\n")[1:]
	stringContent = strings.Join(element, "\n")
	bytes := []byte(stringContent)

	file, err := os.Create("test/.gitpot/index")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()
	file.Write(bytes)
}