package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/bucketSort"
)

func main() {
	nums := []int{0, 1, 2, 2, 0, 0, 1, 1, 1, 0, 0}
	fmt.Println(bucketSort.BucketSort(nums))
}
