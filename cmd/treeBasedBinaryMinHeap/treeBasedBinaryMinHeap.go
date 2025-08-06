package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/treeBasedBinaryMinHeap"
)

func main() {
	// 4. Tree-Based Min Heap
	fmt.Println("\n=== Tree-Based Min Heap ===")
	treeMinHeap := treeBasedBinaryMinHeap.NewMinHeap()
	treeMinHeap.Insert(10)
	treeMinHeap.Insert(20)
	treeMinHeap.Insert(15)
	treeMinHeap.Insert(5)
	fmt.Printf("Size: %d\n", treeMinHeap.Size()) // Size: 4
	if val, ok := treeMinHeap.Peek(); ok {
		fmt.Printf("Peek: %d\n", val) // Peek: 5
	}
	if val, ok := treeMinHeap.ExtractMin(); ok {
		fmt.Printf("Extracted Min: %d\n", val) // Extracted Min: 5
	}
	fmt.Printf("Size after extraction: %d\n", treeMinHeap.Size()) // Size: 3
}
