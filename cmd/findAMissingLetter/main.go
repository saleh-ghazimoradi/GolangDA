package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/findAMissingLetter"
)

func main() {
	input := "the quick brown box jumps over a lazy dog"
	result := findAMissingLetter.FindMissingLetter(input)
	fmt.Printf("Missing letter in %q: %c\n", input, result)
}
