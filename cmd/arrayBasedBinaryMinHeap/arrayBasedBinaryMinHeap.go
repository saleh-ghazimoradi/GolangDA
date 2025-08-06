package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/arrayBasedBinaryMinHeap"
)

func main() {
	// 3. Array-Based Min Heap
	fmt.Println("\n=== Array-Based Min Heap ===")
	arrayMinHeap := arrayBasedBinaryMinHeap.NewMinHeap()
	arrayMinHeap.Insert(10)
	arrayMinHeap.Insert(20)
	arrayMinHeap.Insert(15)
	arrayMinHeap.Insert(5)
	fmt.Printf("Size: %d\n", arrayMinHeap.Size())
	if val, ok := arrayMinHeap.Peek(); ok {
		fmt.Printf("Peek: %d\n", val)
	}
	if val, ok := arrayMinHeap.ExtractMin(); ok {
		fmt.Printf("Extracted Min: %d\n", val)
	}
	fmt.Printf("Size after extraction: %d\n", arrayMinHeap.Size())
}
