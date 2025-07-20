package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/selectionSort"
)

func main() {
	nums := []int{88, 23, 15, 70, 33}
	fmt.Println(nums)
	fmt.Println(selectionSort.SelectionSort(nums))
}
