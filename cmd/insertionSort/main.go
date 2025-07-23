package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/insertionSort"
)

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original array:", arr)
	sorted := insertionSort.InsertionSort(arr)
	fmt.Println("Sorted array:", sorted)
}
