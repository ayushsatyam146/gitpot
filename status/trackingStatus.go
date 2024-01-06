package status

import (
	"fmt"

	file "github.com/ayushsatyam146/gitpot/files"
	"github.com/ayushsatyam146/gitpot/index"
)


func UpdateTrackStatusOfWorkingTree(workingTree *Tree, indexTree *Tree) {
	workingTreeMap := make(map[string]*Tree)
	indexTreeMap := make(map[string]*Tree)
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

func PrintTrackStatus(tree *Tree) {
	
	untrackedFiles := []string{}
	PrintTreeTrackStatus(tree, "\t", &untrackedFiles)
	if len(untrackedFiles) != 0 {
		red := "\033[91m"  
		// green := "\033[92m"    
		reset := "\033[0m"
		message := "Untracked files:\n (use \"gitpot add <file>...\" to include in what will be committed)\n"
		fmt.Println(message)
		for file := range untrackedFiles {
			fmt.Println(red + untrackedFiles[file]+ reset)
		}
	}	
}

func PrintTreeTrackStatus(tree *Tree, prefix string, result *[]string) {
	for _, child := range tree.Children {
		if child.IsDir {
			if !child.Tracked {
				text := prefix + "/" + child.Name
				*result = append(*result, text)
			} else {
				PrintTreeTrackStatus(child, prefix + "/" + child.Name,result)
			}
		} else {
			if !child.Tracked {
				text := prefix + "/" + child.Name
				*result = append(*result, text)
			} 
		}
	}
}

func PrintUntrackedFiles() {
	workingTree := convertTrees(file.GetRelTreeFromWorkingDir("test"))
	indexTree := convertTrees(index.GetTreeFromIndex())
	UpdateTrackStatusOfWorkingTree(workingTree, indexTree)
	PrintTrackStatus(workingTree)
}