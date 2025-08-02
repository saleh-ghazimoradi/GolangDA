package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/greatesProductOfAnyThreeNumbers"
)

func main() {
	nums := []int{1, 5, 3, 8, 2, 6, 7}
	fmt.Println("Original array is", nums)
	result := greatesProductOfAnyThreeNumbers.MaxProductOfThree(nums)
	fmt.Printf("Greatest product of three numbers: %d\n", result)
}
