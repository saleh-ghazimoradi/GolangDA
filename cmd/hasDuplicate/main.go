package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/hasDuplicate"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(hasDuplicate.HasDuplicate(nums))
	fmt.Println(hasDuplicate.EfficientHasDuplicate(nums))
	nums = []int{1, 2, 3, 4, 5, 1}
	fmt.Println(hasDuplicate.HasDuplicate(nums))
	fmt.Println(hasDuplicate.EfficientHasDuplicate(nums))

}
