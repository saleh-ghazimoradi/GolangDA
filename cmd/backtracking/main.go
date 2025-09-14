package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/backtracking"
)

func main() {
	root := backtracking.NewNode(1)
	root.Left = backtracking.NewNode(2)
	root.Right = backtracking.NewNode(3)
	root.Left.Left = backtracking.NewNode(0)
	root.Right.Right = backtracking.NewNode(4)
	result := backtracking.HasValidPath(root)
	fmt.Printf("Does a valid path exist? %v\n", result)
}
