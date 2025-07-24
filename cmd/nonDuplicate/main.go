package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/nonDuplicate"
)

func main() {
	input := "minimum"
	result := nonDuplicate.NonDuplicate(input)
	if result != 0 {
		fmt.Printf("First non-duplicated character in %q: %c\n", input, result)
	} else {
		fmt.Printf("No non-duplicated character in %q\n", input)
	}
}
