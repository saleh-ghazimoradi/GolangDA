package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/missingNumber"
)

func main() {
	nums1 := []int{5, 2, 4, 1, 0}
	nums2 := []int{9, 3, 2, 5, 6, 7, 1, 0, 4}
	fmt.Println("nums1: ", nums1)
	fmt.Printf("Missing number: %d\n", missingnumber.MissingNumber(nums1))
	fmt.Println("nums2: ", nums2)
	fmt.Printf("Missing number: %d\n", missingnumber.MissingNumber(nums2))
}
