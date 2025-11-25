package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/selectionSort"
)

func main() {
	nums := []int{9, 4, 6, 1, 2, 8}
	fmt.Println(nums)
	fmt.Println(selectionSort.SelectionSort(nums))
}
