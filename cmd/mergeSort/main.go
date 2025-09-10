package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/mergeSort"
)

func main() {
	nums := []int{5, 2, 3, 9, 1, 4, 7, 6, 8, 10}
	fmt.Println(mergeSort.MergeSort(nums))
}
