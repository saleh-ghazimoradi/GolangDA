package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/binaryTree"
)

func main() {
	tree := &binaryTree.BinaryTree{}
	values := []int{5, 3, 7, 1, 4, 6, 8}
	for _, v := range values {
		tree.Insert(v)
	}

	fmt.Println("In-order traversal: ")
	tree.InOrderTraversal()
	fmt.Println("Search for 4:", tree.Search(4))
	fmt.Println("Search for 9:", tree.Search(9))
}
