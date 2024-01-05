package index

import (
	file "github.com/ayushsatyam146/gitpot/files"
	"github.com/ayushsatyam146/gitpot/utils"
)

func GetTreeFromIndex() *file.Tree{
	indexFilePath := "test/.gitpot/index"
	indexFileContent := file.ReadFile(indexFilePath)
	indexFileContent = []byte("tree\n" + string(indexFileContent))
	hash,_ := utils.GetSHA1(indexFileContent)
	// fmt.Println(hash)
	// fmt.Println(string(indexFileContent))
	tree := file.GetTreeFromHash("test/.gitpot", hash, "root")
	// file.PrintTree(tree)
	return tree
}