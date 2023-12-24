package main

import (
	"fmt"

	file "github.com/ayushsatyam146/gitpot/files"
)

func main() {
	// cmd.Execute()
	tree := file.GetTreeFromDir("test")
	file.PrintTree(tree)
	hash := file.WriteToGitpotFromTree(tree, "test/.gitpot")
	fmt.Println(hash)
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
