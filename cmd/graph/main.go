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

	// 2. Hash table based graph implementation
	h := graph.NewHGraph()

	h.AddNode("A")
	h.AddNode("B")
	h.AddNode("C")
	h.AddEdge("A", "B")
	h.AddEdge("A", "C")
	h.AddEdge("B", "C")
	h.PrintGraph()
}
