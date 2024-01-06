package status

import (
	file "github.com/ayushsatyam146/gitpot/files"
)

func convertTrees(inputTree *file.Tree) *Tree {
	outputTree := &Tree{
		Name: inputTree.Name,
		IsDir: inputTree.IsDir,
		Value: inputTree.Value,
		Children: []*Tree{},
	}
	for _, child := range inputTree.Children {
		outputTree.Children = append(outputTree.Children, convertTrees(child))
	}
	return outputTree
}