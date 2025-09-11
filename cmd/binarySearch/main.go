package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/binarySearch"
)

func main() {
	arr := []int{2, 3, 4, 10, 40, 50, 60, 70}
	target := 10
	result := binarySearch.BinarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Element %d found at index %d\n", target, result)
	} else {
		fmt.Printf("Element %d not found\n", target)
	}
}
