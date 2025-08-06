package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/arrayBasedBinaryMaxHeap"
)

func main() {
	fmt.Println("=== Array-Based Max Heap ===")
	arrayMaxHeap := arrayBasedBinaryMaxHeap.NewMaxHeap()
	arrayMaxHeap.Insert(10)
	arrayMaxHeap.Insert(20)
	arrayMaxHeap.Insert(15)
	arrayMaxHeap.Insert(30)
	fmt.Printf("Size: %d\n", arrayMaxHeap.Size())
	if val, ok := arrayMaxHeap.Peek(); ok {
		fmt.Printf("Peek: %d\n", val)
	}
	if val, ok := arrayMaxHeap.ExtractMax(); ok {
		fmt.Printf("Extracted Max: %d\n", val)
	}
	fmt.Printf("Size after extraction: %d\n", arrayMaxHeap.Size())
}
