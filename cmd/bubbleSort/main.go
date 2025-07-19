package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/bubbleSort"
)

func main() {
	arr := []int{6, 2, 3, 5, 1, 7}
	fmt.Println(arr)
	fmt.Println(bubbleSort.BubbleSort(arr))
}
