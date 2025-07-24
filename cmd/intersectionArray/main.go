package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/intersectionArray"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{0, 2, 4, 6, 8}
	fmt.Printf("Intersection of %v and %v: %v\n", slice1, slice2, intersectionArray.IntersectionArray(slice1, slice2))
}
