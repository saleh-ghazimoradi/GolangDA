package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/treeNode"
)

func main() {
	root := treeNode.NewTreeNode(1)
	child1 := treeNode.NewTreeNode(2)
	child2 := treeNode.NewTreeNode(3)
	root.AddChild(child1)
	root.AddChild(child2)
	fmt.Println(root)
}
