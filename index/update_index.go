package index

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	file "github.com/ayushsatyam146/gitpot/files"
)

func GetAbsTreeFromPath(dirPath string) *file.Tree {

	elements := strings.Split(dirPath, "/")
	lastElement := elements[len(elements)-1]

	tree := file.Tree{Name: lastElement, IsDir: true}
	children := []*file.Tree{}

	dir, err := os.Open(dirPath)
	if err != nil {
		fmt.Println("Error opening directory:", err)
	}
	defer dir.Close()

	entries, err := dir.ReadDir(-1)
	if err != nil {
		fmt.Println("Error reading directory:", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if entry.Name() == ".git" || entry.Name() == ".gitpot" {
				continue
			}
			children = append(children, GetAbsTreeFromPath(dirPath+"/"+entry.Name()))
		} else {
			bytes := file.ReadFile(dirPath + "/" + entry.Name())
			children = append(children, &file.Tree{Name: entry.Name(), IsDir: false, Value: bytes})
		}
	}
	tree.Children = children
	return &tree
}


func BuildTree(paths []string) *Tree {

	//make this root from index file and if index not present then only create this
	root := &file.Tree{Name: "lamo", IsDir: true}

	for _, path := range paths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				panic("File or directory does not exist")
			} 
		}

		// file or dir make a final leaf node of each and give it as children
		if fileInfo.Mode().IsRegular(){
			absolute_path,_ := filepath.Abs(path)
			elements := strings.Split(path, "/")
			fileName := elements[len(elements)-1]
			elements = elements[:len(elements)-1]
			auxillary_path := strings.Join(elements, "/")
			filePath := file.PathFromRepoRoot(auxillary_path)
			content := file.ReadFile(absolute_path)	
			node := &file.Tree{Name: fileName, IsDir: false, Value: content}
			addPathToTree(root, strings.Split(filePath, "/")[:], node)
			fmt.Println(node)
			PrintTree(node)
		} else if fileInfo.Mode().IsDir() {
			
			dirPath := file.PathFromRepoRoot(path)
			// fmt.Println(dirPath)
			tree := GetAbsTreeFromPath(path)
			// PrintTree(tree)
			PrintTree(root)
			addPathToTree(root, strings.Split(dirPath, "/")[:], tree)
			PrintTree(tree)
		} else {
			panic("It's neither a file nor a directory")
		}
	}

	// PrintTree(root)
	return root
}

func addPathToTree(root *file.Tree, path []string, subtree *file.Tree) {
	current := root
	for i, segment := range path {

		// PrintTree(root)
		var child *file.Tree
		// Check if the segment already exists in children
		for _, existingChild := range current.Children {
			if existingChild.Name == segment {
				child = existingChild
				break
			}
		}

		// If not, create a new child
		if child == nil {
			isDir := i < len(path)-1 // Check if it's a directory
			child = &file.Tree{Name: segment, IsDir: isDir}
			current.Children = append(current.Children, child )
		}

		current = child
	}
	current.Children = append(current.Children, subtree.Children...)
	// PrintTree(root)
}

func PrintTree(tree *file.Tree) {
	printTreeHandler(tree, 0)
}

func printTreeHandler(tree *file.Tree, level int) {
	for i, child := range tree.Children {

		temp := ""
		for i := 0; i < level; i++ {
			temp += "│    "
		}
		fmt.Print(temp)

		if i == len(tree.Children)-1 {
			fmt.Print("└───")
		} else {
			fmt.Print("├───")
		}
		if child.IsDir {
			fmt.Println(child.Name)
			printTreeHandler(child, level+1)
		} else {
			fmt.Println(child.Name)
		}
	}
}
