package index

import (
	"os"
	"path/filepath"
	"strings"

	file "github.com/ayushsatyam146/gitpot/files"
)

func BuildTree(paths []string) *file.Tree {
	root := &file.Tree{Name: "root", IsDir: true}

	for _, path := range paths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				panic("File or directory does not exist")
			} 
		}

		if fileInfo.Mode().IsRegular(){
			absolute_path,_ := filepath.Abs(path)
			elements := strings.Split(path, "/")
			fileName := elements[len(elements)-1]
			elements = elements[:len(elements)-1]
			auxillary_path := strings.Join(elements, "/")
			filePath := file.PathFromRepoRoot(auxillary_path)
			content := file.ReadFile(absolute_path)	
			node := &file.Tree{Name: fileName, IsDir: false, Value: content}
			addPathToTree(root, strings.Split(filePath, "/")[:], node, false)
		} else if fileInfo.Mode().IsDir() {
			
			dirPath := file.PathFromRepoRoot(path)
			tree := file.GetAbsTreeFromPath(path)
			addPathToTree(root, strings.Split(dirPath, "/")[:], tree, true)
		} else {
			panic("It's neither a file nor a directory")
		}
	}

	return root
}

func addPathToTree(root *file.Tree, path []string, subtree *file.Tree, isDir bool) {
	current := root
	for _, segment := range path {
		var child *file.Tree
		for _, existingChild := range current.Children {
			if existingChild.Name == segment {
				child = existingChild
				break
			}
		}
		if child == nil {
			child = &file.Tree{Name: segment, IsDir: true}
			current.Children = append(current.Children, child )
		}
		current = child
	}
	if isDir {
		current.Children = append(current.Children, subtree.Children...)
	} else {
		current.Children = append(current.Children, subtree)
	}
}