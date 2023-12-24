package files

import (
	"os"
)

func WriteTreeToDisk(tree *Tree, path string)  {

	os.MkdirAll(path, 0755)

	for _, child := range tree.children {
		if !child.isDir {
			// extract the values from the child and write to disk
			file, err := os.Create(path + "/" + child.name)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			file.Write(child.value)
		} else {
			WriteTreeToDisk(child, path + "/" + child.name)
		}
	}

}
