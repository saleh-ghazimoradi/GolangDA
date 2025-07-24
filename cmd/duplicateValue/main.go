package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/duplicateValue"
)

func main() {
	input := []string{"a", "b", "c", "d", "c", "e", "f"}
	result := duplicateValue.DuplicateValue(input)
	fmt.Printf("First duplicate in %v: %s\n", input, result)
}
