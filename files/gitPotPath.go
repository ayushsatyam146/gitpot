package files

import (
	"os"
)

func GitPotDir(path string) string {
	potentialGitPotDir := path + "/.gitpot"
	// check if path exists and is a directory

	dirInfo, err := os.Stat(potentialGitPotDir)
	if err == nil {
		if dirInfo.IsDir() {
			return potentialGitPotDir
		} else {
			panic("The given gitpot path is not a directory")
		}
	} else {
		if os.IsNotExist(err) {
			return GitPotDir(path + "/..")
		}
	}
	panic("err")
}

// give "" in path to get the gitpot directory
func GitPotPath(path string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return GitPotDir(currentDir) + "/" + path
}


func WorkingCopyPath(path string) string {
	return GitPotPath("") + "../" + path
}