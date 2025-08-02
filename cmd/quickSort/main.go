package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/quickSort"
)

func main() {
	nums := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original array:", nums)
	sorted := quickSort.QuickSort(nums)
	fmt.Println("Sorted array:", sorted)
}
