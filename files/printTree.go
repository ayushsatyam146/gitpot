package files

import (
	"fmt"
)

func PrintTree(tree *Tree) {
	printTreeHandler(tree, 0)
}

func printTreeHandler(tree *Tree, level int) {
	for i, child := range tree.children {

		temp := ""
		for i := 0; i < level; i++ {
			temp += "│    "
		}
		fmt.Print(temp)

		if i == len(tree.children)-1 {
			fmt.Print("└───")
		} else {
			fmt.Print("├───")
		}
		if child.isDir {
			fmt.Println(child.name)
			printTreeHandler(child, level+1)
		} else {
			fmt.Println(child.name)
		}
	}
}
