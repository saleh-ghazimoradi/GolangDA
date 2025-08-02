package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/maxNumber"
)

func main() {
	nums := []int{5, 2, 9, 1, 7, 6, 3}
	fmt.Println("Original array:", nums)

	fmt.Printf("Max (O(N²)): %d\n", maxNumber.QuadraticMaxNumber(nums))
	fmt.Printf("Max (O(N log N)): %d\n", maxNumber.MaxNumber(nums))
	fmt.Printf("Max (O(N)): %d\n", maxNumber.LinearMaxNumber(nums))
}
