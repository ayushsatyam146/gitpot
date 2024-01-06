package index

import (
	"os"
	"path/filepath"
	"strings"

	file "github.com/ayushsatyam146/gitpot/files"
)

func BuildTree(paths []string) *file.Tree {

	// ------------------will make git add more robust------------------
	// get this root after parsing any index file if it exists, if not then create a new one
	// read index file and compute the hash and then use GetTreeFromHash to get the tree
	// for comparing any new path or file that is added via `add` command this same tree will be used for comparison

	// ------------------will tell about tracked/untracked files------------------
	// we can also check where exactly is this generated tree different 
	// from the tree that we get from GetTreeFromDir by passing working dir in this function

	// to know about staged and unstaged change we can just get the tree 'root' in a similar manner (mentioned above to above)
  // from the current index file. After that in that tree we can just compute the hashes of each file from the working dir mentioned in the 'root' tree
	// if hash is same we can say it doesn't have any changes if it's not same we can say we have unstaged changes
	// after that upon `gitpot add` we can just update index from the current working dir and get the index updated to latest possible version

	root := &file.Tree{Name: "root", IsDir: true}
	// indexContent := string(file.ReadFile("test/.gitpot/index"))
	// if indexContent != "" {
	// 	root = GetTreeFromIndex()
	// }

	for _, path := range paths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				panic("File or directory does not exist")
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
		current.Children = append(current.Children, subtree.Children...)
	} else {
		for i, child := range current.Children {
			if child.Name == subtree.Name {
				current.Children[i] = subtree
				break
			}
		}
		current.Children = append(current.Children, subtree)
	}
}
