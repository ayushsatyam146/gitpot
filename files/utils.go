package files

import (
	"os"
	"path/filepath"
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

func pathFromRepoRoot(path string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	destPath := currentDir + "/" + path
	relativePath, err := filepath.Rel(WorkingCopyPath(""), destPath)
	if err != nil {
		panic(err)
	}
	return relativePath
}
