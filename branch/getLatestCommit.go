package branch

import (
	file "github.com/ayushsatyam146/gitpot/files"
)

func GetLatestCommit() string {
	branch := file.ReadFile("test/.gitpot/HEAD")
	latestCommit := file.ReadFile("test/.gitpot/" + string(branch))
	return string(latestCommit)
}