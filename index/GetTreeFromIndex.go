package index

import (
	"strings"

	file "github.com/ayushsatyam146/gitpot/files"
	"github.com/ayushsatyam146/gitpot/utils"
)

func GetTreeFromIndex() *file.Tree{
	indexFilePath := "test/.gitpot/index"
	indexFileContent := file.ReadFile(indexFilePath)
	if len(string(indexFileContent)) == 0 {
		return &file.Tree{Name: "root", IsDir: true, Children: []*file.Tree{}}
	}
	indexFileContent = []byte("tree\n" + string(indexFileContent))
	hash,_ := utils.GetSHA1(indexFileContent)
	tree := file.GetTreeFromHash("test/.gitpot", hash, "root")
	return tree
}

func GetTreeFromPrevCommit() *file.Tree{
	commitFile:= string(file.ReadFile("test/.gitpot/HEAD"))
	commitHash := string(file.ReadFile("test/.gitpot/" + commitFile))
	if len(commitHash) == 0 {
		return &file.Tree{Name: "root", IsDir: true, Children: []*file.Tree{}}
	}
	commitContent := string(file.ReadFile("test/.gitpot/objects/" + commitHash[:2] + "/" + commitHash[2:]))
	treeRow := strings.Split(commitContent, "\n")[1]
	treeHash := strings.Split(treeRow, " ")[1]
	tree := file.GetTreeFromHash("test/.gitpot", treeHash, "root")
	return tree
}