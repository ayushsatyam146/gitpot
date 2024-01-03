package files

import (
	"os"
	"path/filepath"
)

func GitPotDir(path string) string {
	potentialGitPotDir := path + "/.gitpot"

	dirInfo, err := os.Stat(potentialGitPotDir)
	if err == nil {
		if dirInfo.IsDir() {
			return filepath.Clean(potentialGitPotDir)
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

func GitPotPath(path string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return GitPotDir(currentDir) + "/" + path
}

func WorkingCopyPath(path string) string {
	return filepath.Clean(GitPotDir(path) + "/../")
}
