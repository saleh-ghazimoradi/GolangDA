package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/greatesValue"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(greatesValue.GreatestValue(nums...))
}
