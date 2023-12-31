package files

import (
	"os"
	"path/filepath"
	"strings"

	utils "github.com/ayushsatyam146/gitpot/utils"
)

func inRepo() bool {
	res := GitPotPath("")
	if res != "" {
		return true
	}
	return false
}

func assertInRepo() {
	if !inRepo() {
		panic("Not in a gitpot repository")
	}
}

func PathFromRepoRoot(path string) string {
	abs1, _ := filepath.Abs(WorkingCopyPath(path))
	abs2, _ := filepath.Abs(path)
	rel, _ := filepath.Rel(abs1, abs2)
	// fmt.Println(abs1, abs2, rel)
	// fmt.Println(WorkingCopyPath(path))
	return rel
}


func WriteTreeToWorkingDir(tree *Tree, path string)  {
	os.MkdirAll(path, 0755)

	for _, child := range tree.children {
		if !child.isDir {
			// extract the values from the child and write to disk
			file, err := os.Create(path + "/" + child.name)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			file.Write(child.value)
		} else {
			WriteTreeToWorkingDir(child, path + "/" + child.name)
		}
	}
}

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