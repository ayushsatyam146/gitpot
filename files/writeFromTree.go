package files

import (
	utils "github.com/ayushsatyam146/gitpot/utils"
)

func WriteToGitpotFromTree(tree *Tree, path string) string {

	TreeFileContent := "tree\n"
	for _, child := range tree.children {
		if !child.isDir {
			utils.WriteToObjectsDir(path, child.value, false)
			hash, _ := utils.GetSHA1(child.value)
			TreeFileContent += "blob " + hash + " " + child.name + "\n"
		} else {
			hash := WriteToGitpotFromTree(child, path)
			TreeFileContent += "tree " + hash + " " + child.name + "\n"
		}
	}

	utils.WriteToObjectsDir(path, []byte(TreeFileContent), true)
	hash, _ := utils.GetSHA1([]byte(TreeFileContent))
	return hash

}
