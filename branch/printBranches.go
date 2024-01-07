package branch

import (
	"fmt"
	"os"
)

func PrintBrancheNames() {
	dirPath := "test/.gitpot/refs/heads"
	dir, err := os.Open(dirPath)
	if err != nil {
		fmt.Println("Error opening directory:", err)
	}
	defer dir.Close()

	files, err := dir.ReadDir(-1)
	if err != nil {
		fmt.Println("Error reading directory:", err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}