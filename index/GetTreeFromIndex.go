package index

import (
	"strings"

	file "github.com/ayushsatyam146/gitpot/files"
	"github.com/ayushsatyam146/gitpot/utils"
)

func GetTreeFromIndex() *file.Tree{
	indexFilePath := "test/.gitpot/index"
	indexFileContent := file.ReadFile(indexFilePath)
	indexFileContent = []byte("tree\n" + string(indexFileContent))
	hash,_ := utils.GetSHA1(indexFileContent)
	tree := file.GetTreeFromHash("test/.gitpot", hash, "root")
	return tree
}

func GetTreeFromPrevCommit() *file.Tree{
	commitHash := string(file.ReadFile("test/.gitpot/HEAD"))
	commitContent := string(file.ReadFile("test/.gitpot/objects/" + commitHash[:2] + "/" + commitHash[2:]))
	treeRow := strings.Split(commitContent, "\n")[1]
	treeHash := strings.Split(treeRow, " ")[1]
	tree := file.GetTreeFromHash("test/.gitpot", treeHash, "root")
	return tree
}