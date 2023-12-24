package main

import (
	"fmt"

	// "github.com/ayushsatyam146/gitpot/cmd"
	file "github.com/ayushsatyam146/gitpot/files"
)

func main() {
	// cmd.Execute()
	tree := file.GetTreeFromDir("test")
	file.PrintTree(tree)
	hash := file.WriteTreeToGitpot(tree, "test/.gitpot")
	fmt.Println(hash)
	tree2 := file.GetTreeFromHash("test/.gitpot", hash, "")
	file.PrintTree(tree2)

	// res, _ := files.ListFilesRecursive(".")
	// for _, file := range res {
	// 	fmt.Println(file)
	// }
	// file.RemoveEmptyDirs(".")
	// file.Write("./cmd/lol.txt","doggy bhao \n haa bao")
	// print(file.ReadFileToString("./a.txt"))
	// cmd.Execute()
	// str := file.GitPotPath("HEAD")
	// fmt.Println(str)
	// fmt.Println(file.WorkingCopyPath("files"))
}
