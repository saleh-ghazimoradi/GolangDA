package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/trie"
)

func main() {
	t := trie.NewTrie()
	t.Insert("apple")
	fmt.Println(t.Search("apple"))   // true
	fmt.Println(t.Search("app"))     // false
	fmt.Println(t.StartsWith("app")) // true
	t.Insert("app")
	fmt.Println(t.Search("app")) // true
}
