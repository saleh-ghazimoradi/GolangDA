package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/graph"
)

func main() {
	// 1. Object-oriented based graph implementation
	g := graph.NewGraph()
	g.AddNode("A")
	g.AddNode("B")
	g.AddNode("C")
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "C")
	g.PrintGraph()
	fmt.Println()
}
