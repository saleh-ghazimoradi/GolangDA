package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/quickSelect"
)

func main() {
	nums := []int{64, 34, 25, 12, 22, 11, 90}
	k := 3
	fmt.Println("Original array:", nums)

	result := quickSelect.QuickSelect(nums, k)
	fmt.Printf("The %dth smallest element is: %d\n", k, result)
}
