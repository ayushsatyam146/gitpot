package index

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	file "github.com/ayushsatyam146/gitpot/files"
)

func BuildTree(paths []string) *file.Tree {
	root := &file.Tree{Name: "root", IsDir: true}
	indexContent := string(file.ReadFile("test/.gitpot/index"))
	if indexContent != "" {
		root = GetTreeFromIndex()
	}

	for _, path := range paths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("File or directory does not exist")
			}
		}

		if fileInfo.Mode().IsRegular() {
			absolute_path, _ := filepath.Abs(path)
			elements := strings.Split(path, "/")
			fileName := elements[len(elements)-1]
			elements = elements[:len(elements)-1]
			auxillary_path := strings.Join(elements, "/")
			filePath := file.PathFromRepoRoot(auxillary_path)
			content := file.ReadFile(absolute_path)
			node := &file.Tree{Name: fileName, IsDir: false, Value: content}
			if filePath == "." {
				addPathToTree(root, []string{}, node, false)
			} else {
				addPathToTree(root, strings.Split(filePath, "/")[:], node, false)
			}
		} else if fileInfo.Mode().IsDir() {

			dirPath := file.PathFromRepoRoot(path)
			if dirPath == "." {
				dirPath = ""
			}
			tree := file.GetAbsTreeFromPath(path)
			if dirPath == "." || dirPath == "" {
				addPathToTree(root, []string{}, tree, true)
			} else {
				addPathToTree(root, strings.Split(dirPath, "/")[:], tree, true)
			}
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
			current.Children = append(current.Children, child)
		}
		current = child
	}
	if isDir {
		// current.Children = append(current.Children, subtree.Children...)
		childrenMap := make(map[string]*file.Tree)
		for _, child := range current.Children {
			childrenMap[child.Name] = child
		}
		for _, child := range subtree.Children {
			if _, ok := childrenMap[child.Name]; ok {
				childrenMap[child.Name] = child
			} else {
				childrenMap[child.Name] = child
			}
		}
		current.Children = []*file.Tree{}
		for _, child := range childrenMap {
			current.Children = append(current.Children, child)
		}
	} else {
		if len(current.Children) == 0 {
			current.Children = append(current.Children, subtree)
			return
		}
		fileWritten := false
		for i, child := range current.Children {
			if child.Name == subtree.Name {
				current.Children[i] = subtree
				fileWritten = true
				break
			}
		}
		if !fileWritten {
			current.Children = append(current.Children, subtree)
		}
	}
}
