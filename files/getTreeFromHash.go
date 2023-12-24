package files

import (
	"io"
	"os"
	"strings"
)

// path will be "test/.gitpot" i.e. the path of the gitpot directory hash is just the hash of the tree object
func GetTreeFromHash(path string, hash string, treeName string) *Tree {

	filePath := path + "/objects/" + hash[:2] + "/" + hash[2:]

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// content from the tree hash file
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	tree := Tree{name: treeName, isDir: true}
	children := []*Tree{}

	hashEntries := strings.Split(string(content), "\n")
	for _, hashEntry := range hashEntries {
		elements := strings.Split(hashEntry, " ")
		if len(elements) == 1 {
			continue
		}
		hash := elements[1]
		name := elements[2]
		if elements[0] == "tree" {
			children = append(children, GetTreeFromHash(path, hash, name))
		} else {

			file, err := os.Open(path + "/objects/" + hash[:2] + "/" + hash[2:])
			if err != nil {
				panic(err)
			}
			defer file.Close()

			stat, err := file.Stat()
			if err != nil {
				panic(err)
			}

			fileSize := stat.Size()
			content := make([]byte, fileSize)
			_, err = file.Read(content)
			if err != nil {
				panic(err)
			}

			children = append(children, &Tree{name: name, isDir: false, value: content})
		}
	}
	tree.children = children
	return &tree
}
