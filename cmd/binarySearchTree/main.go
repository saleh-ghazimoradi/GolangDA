package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/binarySearchTree"
)

func main() {
	tree := &binarySearchTree.BinarySearchTree{}

	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, v := range values {
		tree.Insert(v)
	}

	// Test traversals
	fmt.Println("In-order traversal:")
	tree.InOrderTraversal() // Expected: 20 30 40 50 60 70 80

	fmt.Println("Pre-order traversal:")
	tree.PreOrderTraversal() // Expected: 50 30 20 40 70 60 80

	fmt.Println("Post-order traversal:")
	tree.PostOrderTraversal() // Expected: 20 40 30 60 80 70 50

	// Test search
	fmt.Println("Search for 40:", tree.Search(40)) // Expected: true
	fmt.Println("Search for 90:", tree.Search(90)) // Expected: false

	// Test deletion
	fmt.Println("Deleting 30...")
	tree.Delete(30)
	fmt.Println("In-order traversal after deletion:")
	tree.InOrderTraversal() // Expected: 20 40 50 60 70 80
}
