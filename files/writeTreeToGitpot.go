package files

import (
	"strings"

	utils "github.com/ayushsatyam146/gitpot/utils"
)

func WriteTreeToGitpot(tree *Tree, path string) string {

	TreeFileContent := "tree\n"
	for _, child := range tree.children {
		if !child.isDir {
			utils.WriteToObjectsDir(path, child.value, false)
			hash, _ := utils.GetSHA1(child.value)
			TreeFileContent += "blob " + hash + " " + child.name + "\n"
		} else {
			hash := WriteTreeToGitpot(child, path)
			elements := strings.Split(child.name, "/")
			lastElement := elements[len(elements)-1]
			// fmt.Println(lastElement)
			TreeFileContent += "tree " + hash + " " + lastElement + "\n"
		}
	}

	utils.WriteToObjectsDir(path, []byte(TreeFileContent), true)
	hash, _ := utils.GetSHA1([]byte(TreeFileContent))
	return hash

}
