package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/stack"
)

func main() {
	s := stack.NewStack[int]()
	fmt.Println("Stack Operations:")
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Printf("Stack size: %d\n", s.Size())

	if item, ok := s.Peek(); ok {
		fmt.Printf("Stack peek: %d\n", item)
	}

	for !s.IsEmpty() {
		if item, ok := s.Pop(); ok {
			fmt.Printf("Popped: %d\n", item)
		}
	}
	fmt.Printf("Stack is empty: %v\n", s.IsEmpty())
}
