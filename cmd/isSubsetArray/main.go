package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/isSubsetArray"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(isSubsetArray.IsSubsetArray(slice1, slice2))
	fmt.Println(isSubsetArray.MoreEfficientIsSubsetArray(slice1, slice2))
}
