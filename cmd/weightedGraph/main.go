package main

import (
	"github.com/saleh-ghazimoradi/GolangDA/weightedGraph"
)

func main() {
	g := weightedGraph.NewGraph()
	g.AddNode("A")
	g.AddNode("B")
	g.AddNode("C")
	g.AddEdge("A", "B", 5)
	g.AddEdge("B", "C", 3)
	g.PrintGraph()
}
