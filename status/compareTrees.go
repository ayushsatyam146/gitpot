package status

import (
	file "github.com/ayushsatyam146/gitpot/files"
	// "github.com/ayushsatyam146/gitpot/index"
	// "github.com/ayushsatyam146/gitpot/utils"
)

// func CompareTrees(workingTree *file.Tree, indexTree *file.Tree) {
// 	workingTreeChildren := []string{}
// 	for _, child := range workingTree.Children {
// 		workingTreeChildren = append(workingTreeChildren, child.Name)
// 	}
// 	indexTreeChildren := []string{}
// 	for _, child := range indexTree.Children {
// 		indexTreeChildren = append(indexTreeChildren, child.Name)
// 	}
// 	fmt.Println(workingTreeChildren)
// 	fmt.Println(indexTreeChildren)
// }

// CompareTrees compares two trees and returns the nodes present in tree1 but not in tree2.
func CompareTrees(tree1, tree2 *file.Tree) []*file.Tree {
	var diffNodes []*file.Tree

	compareNodes(tree1, tree2, &diffNodes)

	return diffNodes
}

func compareNodes(node1, node2 *file.Tree, diffNodes *[]*file.Tree) {
	if node1 == nil {
		return
	}

	var found bool

	// Search for node1 in tree2
	for _, child2 := range node2.Children {
		if node1.Name == child2.Name && node1.IsDir == child2.IsDir {
			found = true
			compareNodes(node1, child2, diffNodes)
			break
		}
	}

	// If node1 is not found in tree2, add it to the difference list
	if !found {
		*diffNodes = append(*diffNodes, node1)
	}

	// Recursively compare child nodes
	for _, child1 := range node1.Children {
		compareNodes(child1, node2, diffNodes)
	}
}
