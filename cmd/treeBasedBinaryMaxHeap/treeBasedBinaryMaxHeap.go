package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/treeBasedBinaryMaxHeap"
)

func main() {
	fmt.Println("=== Testing Tree-Based Max Heap ===")
	heap := treeBasedBinaryMaxHeap.NewMaxHeap()
	heap.Insert(10)
	heap.Insert(20)
	heap.Insert(15)
	heap.Insert(30)
	fmt.Printf("Size: %d\n", heap.Size())
	if val, ok := heap.Peek(); ok {
		fmt.Printf("Peek: %d\n", val)
	}
	if val, ok := heap.ExtractMax(); ok {
		fmt.Printf("Extracted Max: %d\n", val)
	}
	fmt.Printf("Size after extraction: %d\n", heap.Size())
}
