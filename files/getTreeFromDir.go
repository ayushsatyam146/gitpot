package files

import (
	"fmt"
	"os"
)

// Give dir path and get tree of the directory
func GetTreeFromDir(dirPath string) *Tree {
	tree := Tree{name: dirPath, isDir: true}
	children := [] *Tree{}

	dir, err := os.Open(dirPath)
	if err != nil {
		fmt.Println("Error opening directory:", err)
	}
	defer dir.Close()

	entries, err := dir.ReadDir(-1)
	if err != nil {
		fmt.Println("Error reading directory:", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if(entry.Name() == ".git" || entry.Name() == ".gitpot") {
				continue
			}
			children = append(children, GetTreeFromDir(dirPath + "/" + entry.Name()))
		} else {
			children = append(children, &Tree{name: entry.Name(), isDir: false})
		}
	}
	tree.children = children
	return &tree
}
