package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/orderedArray"
)

func main() {
	arr := orderedArray.NewOrderedArray()
	arr.Add(3)
	arr.Add(1)
	arr.Add(2)
	arr.Add(1)
	fmt.Println("array elements:", arr.Element())
	num, exists := arr.Find(2)
	fmt.Printf("find num: %v and exists: %v\n", num, exists)
	fmt.Println("array size:", arr.Size())
	arr.Remove(1)
	fmt.Println("after removing 1:", arr.Element())
	if val, ok := arr.Get(1); ok {
		fmt.Println("element at index 1", val)
	}
	arr.Clear()
	fmt.Println("after clearing:", arr.Element())
}
