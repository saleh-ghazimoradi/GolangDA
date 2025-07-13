package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/set"
)

func main() {
	s := set.NewSet()
	s.Add(1)
	s.Add(2)
	s.Add(1)
	fmt.Println("set elements:", s.Element())
	fmt.Println("contain 2:", s.Contains(2))
	fmt.Println("size:", s.Size())
	s.Remove(1)
	fmt.Println("after removing 1:", s.Element())
	s.Clear()
	fmt.Println("after clearing:", s.Element())
}
