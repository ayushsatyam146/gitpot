package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// make it get tree from file or dir path
// Give dir path and get tree of the directory
// make it construct tree of nested tree structures as well like docs/something/else/a.txt
// make it smart enough to construct unified tree if we have docs/something/else/a.txt and docs/something/else/b.txt

func GetRelTreeFromWorkingDir(path string) *Tree{

	fmt.Println(filepath.Abs(path))
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			panic("File or directory does not exist")
		} 
	}

	// Check if it's a file
	if fileInfo.Mode().IsRegular() {
		// if a file
		
		elements := strings.Split(path, "/")
		fileName := elements[len(elements)-1]
		elements = elements[:len(elements)-1]
		auxillary_path := strings.Join(elements, "/")
		auxillary_path = PathFromRepoRoot(auxillary_path)
		elements = strings.Split(auxillary_path, "/")

		abs_path,_ := filepath.Abs(path)
		bytes := ReadFile(abs_path)
		
		node := &Tree{Name: fileName, IsDir: false, Value: bytes}

		current_node := &Tree{}
		tree := current_node
		for _,entry := range elements {
			if entry == ".git" || entry == ".gitpot" {
				continue
			}
			new_node := Tree{Name: entry, IsDir: true}
			current_node.Children = append(current_node.Children, &new_node)
			current_node = &new_node
		}

		current_node.Children = append(current_node.Children, node)

		PrintTree(tree)
		return tree

	} else if fileInfo.Mode().IsDir() {
		auxillary_path := PathFromRepoRoot(path)
		current_node := &Tree{}
		tree := current_node

		if auxillary_path == "." {
			return GetAbsTreeFromPath(path)
		}
		elements := strings.Split(auxillary_path, "/")
		
		for _,entry := range elements {
			if entry == ".git" || entry == ".gitpot" {
				continue
			}
			new_node := Tree{Name: entry, IsDir: true}
			current_node.Children = append(current_node.Children, &new_node)
			current_node = &new_node
		}

		tree2 := GetAbsTreeFromPath(path)
		current_node.Children = append(current_node.Children, tree2.Children...)
		return tree
	} else {
		panic("It's neither a file nor a directory")
	}
}

// can be file can be path
func GetAbsTreeFromPath(dirPath string) *Tree {

	elements := strings.Split(dirPath, "/")
	lastElement := elements[len(elements)-1]

	tree := Tree{Name: lastElement, IsDir: true}
	children := []*Tree{}

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
			bytes := ReadFile(dirPath + "/" + entry.Name())
			children = append(children, &Tree{Name: entry.Name(), IsDir: false, Value: bytes})
		}
	}
	tree.Children = children
	return &tree
}
