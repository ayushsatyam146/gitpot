package main

import (
	"github.com/ayushsatyam146/gitpot/cmd"
	// file "github.com/ayushsatyam146/gitpot/files"
	// "github.com/ayushsatyam146/gitpot/index"
)


func main() {
	cmd.Execute()
	// var args []string
	// tree.Name = "test"
	// hash := file.WriteTreeToGitpot(tree, "test/.gitpot")
	// file.PrintTree(tree)
	// fmt.Println(hash)
	// tree2 := file.GetTreeFromHash("test/.gitpot", hash, "")
	// // file.PrintTree(tree2)
	// file.WriteTreeToWorkingDir(tree2, "test-write")
	// c := file.ReadFile("test/.gitpot/index")
	// fmt.Println(string(c))

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
