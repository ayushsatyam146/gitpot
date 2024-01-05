package status

import (
	"fmt"

	file "github.com/ayushsatyam146/gitpot/files"
	// "github.com/ayushsatyam146/gitpot/index"
	// "github.com/ayushsatyam146/gitpot/utils"
)


func UpdateTrackStatusOfWorkingTree(workingTree *file.Tree, indexTree *file.Tree) {
	workingTreeMap := make(map[string]*file.Tree)
	indexTreeMap := make(map[string]*file.Tree)
	for _, child := range workingTree.Children {
		workingTreeMap[child.Name] = child
	}
	for _, child := range indexTree.Children {
		indexTreeMap[child.Name] = child
	}
	for key, value := range workingTreeMap {
		if indexTreeMap[key] == nil {
			value.Tracked = false
		} else {
			if value.IsDir {
				value.Tracked = true
				UpdateTrackStatusOfWorkingTree(value, indexTreeMap[key])
			} else {
				value.Tracked = true
			}
		}
	}
}

func PrintTrackStatus(tree *file.Tree) {
	message := "Untracked files:\n (use \"gitpot add <file>...\" to include in what will be committed)\n"
	fmt.Println(message)
	PrintTreeTrackStatus(tree, "\t")
}

func PrintTreeTrackStatus(tree *file.Tree, prefix string) {
	red := "\033[91m"    
	// green := "\033[92m"  
	reset := "\033[0m"
	for _, child := range tree.Children {
		if child.IsDir {
			if !child.Tracked {
				text := prefix + "/" + child.Name
				fmt.Println(red + text + reset)
			} else {
				PrintTreeTrackStatus(child, prefix + "/" + child.Name)
			}
		} else {
			if !child.Tracked {
				text := prefix + "/" + child.Name
				fmt.Println(red + text+ reset)
			} 
		}
	}
}