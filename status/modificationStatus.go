package status

import (
	"fmt"

	file "github.com/ayushsatyam146/gitpot/files"
	"github.com/ayushsatyam146/gitpot/index"
	"github.com/ayushsatyam146/gitpot/utils"
)

func UpdateModifiedStatusOfWorkingTree(workingTree *Tree, curr_index *Tree, prev_index *Tree) {
	workingTreeMap := make(map[string]*Tree)
	curr_indexMap := make(map[string]*Tree)
	prev_indexMap := make(map[string]*Tree)
	for _, child := range workingTree.Children {
		workingTreeMap[child.Name] = child
	}
	for _, child := range curr_index.Children {
		curr_indexMap[child.Name] = child
	}
	last_commit_hash := string(file.ReadFile("test/.gitpot/refs/heads/master"))
	if(last_commit_hash != "") {
		for _, child := range prev_index.Children {
			prev_indexMap[child.Name] = child
		}
	}
	for key, value := range curr_indexMap {
		if workingTreeMap[key] == nil {
			value.Modified = true
			// throw an exception that file has been deleted from working dir and 'git add' needs to run again
		} else {
			if(len(prev_indexMap) == 0) {
				if !value.IsDir {
					currIndexHash,_ := utils.GetSHA1(value.Value)
					workingTreeHash,_ := utils.GetSHA1(workingTreeMap[key].Value)
					if currIndexHash == workingTreeHash {
						value.Modified = true
						value.Staged = true
					} else {
						value.Modified = true
						value.Staged = false
					}
				} else {
					UpdateModifiedStatusOfWorkingTree(workingTreeMap[key], curr_indexMap[key], prev_indexMap[key])
				}							
			} else {
				if !value.IsDir {
					currIndexHash,_ := utils.GetSHA1(value.Value)
					prevIndexHash := ""
					if prev_indexMap[key] != nil {
						prevIndexHash,_ = utils.GetSHA1(prev_indexMap[key].Value)
					}
					workingTreeHash,_ := utils.GetSHA1(workingTreeMap[key].Value)
					if currIndexHash == workingTreeHash {
						if prevIndexHash != currIndexHash {
							value.Modified = true
							value.Staged = true
						} else {
							value.Modified = false
						}
					} else {
						value.Modified = true
						value.Staged = false
					}
				} else {
					UpdateModifiedStatusOfWorkingTree(workingTreeMap[key], curr_indexMap[key], prev_indexMap[key])
				}
			}
		}
	}
}

func PrintStagedChanges(tree *Tree) {
	for _, child := range tree.Children {
		if child.IsDir {
			PrintStagedChanges(child)
		} else {
			if child.Staged {
				green := "\033[92m"    
				reset := "\033[0m"
				fmt.Println(green + child.Name + reset)
			}
		}
	}
}

func lmao(tree *Tree) {
	for _, child := range tree.Children {
		if child.IsDir {
			lmao(child)
		} else {
			if child.Modified && !child.Staged {
				red := "\033[91m"    
				reset := "\033[0m"
				fmt.Println(red + child.Name + reset)
			}
		}
	}
}

func GetUnstagedChanges(tree *Tree, prefix string, result *[]string) {
	for _, child := range tree.Children {
		if child.IsDir {
			GetUnstagedChanges(child, prefix + "/" + child.Name,result)
		} else {
			if child.Modified && !child.Staged {
				text := prefix + "/" + child.Name
				*result = append(*result, text)
			} 
		}
	}
}

func GetStagedChanges(tree *Tree, prefix string, result *[]string) {
	for _, child := range tree.Children {
		if child.IsDir {
			GetStagedChanges(child, prefix + "/" + child.Name,result)
		} else {
			if child.Modified && child.Staged {
				text := prefix + "/" + child.Name
				*result = append(*result, text)
			} 
		}
	}
}

func PrintModifiedFiles() {
	workingTree := convertTrees(file.GetRelTreeFromWorkingDir("test"))
	curr_indexTree := convertTrees(index.GetTreeFromIndex())
	prev_indexTree := convertTrees(index.GetTreeFromPrevCommit())
	UpdateModifiedStatusOfWorkingTree(workingTree, curr_indexTree, prev_indexTree)
	UnstagedChanges := []string{}
	GetUnstagedChanges(curr_indexTree,"\t", &UnstagedChanges)
	StagedChanges := []string{}
	GetStagedChanges(curr_indexTree,"\t", &StagedChanges)
	if len(UnstagedChanges) != 0  || len(StagedChanges) != 0 {
		fmt.Println("Modified files:")	
	}
	if len(UnstagedChanges) != 0 {
		red := "\033[91m"     
		reset := "\033[0m"
		message := "Changes not staged for commit:\n (use \"gitpot add <file>...\" to update what will be committed)\n"
		fmt.Println(message)
		for file := range UnstagedChanges {
			fmt.Println(red + UnstagedChanges[file]+ reset)
		}
	}
	if len(StagedChanges) != 0 {
		green := "\033[92m"    
		reset := "\033[0m"
		message := "Changes to be committed:\n (use \"gitpot rm --cached <file>...\" to unstage)\n"
		fmt.Println(message)
		for file := range StagedChanges {
			fmt.Println(green + StagedChanges[file]+ reset)
		}
	}
}