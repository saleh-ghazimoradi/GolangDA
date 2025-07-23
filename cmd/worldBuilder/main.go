package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/wordBuilder"
)

func main() {
	letters := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(wordBuilder.WordBuilder(letters))
}
