package files

import (
	"fmt"
)

func PrintTree(tree *Tree) {
	fmt.Println(tree.Name)
	printTreeHandler(tree, 0)
}

func printTreeHandler(tree *Tree, level int) {
	for i, child := range tree.Children {

		temp := ""
		for i := 0; i < level; i++ {
			temp += "│    "
		}
		fmt.Print(temp)

		if i == len(tree.Children)-1 {
			fmt.Print("└───")
		} else {
			fmt.Print("├───")
		}
		if child.IsDir {
			// fmt.Println(child.Name, child.IsDir)
			fmt.Println(child.Name)
			printTreeHandler(child, level+1)
		} else {
			// fmt.Println(child.Name, child.IsDir)
			fmt.Println(child.Name)
		}
	}
}
