package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/graphBFS"
)

func main() {
	g := graphBFS.NewGraph()
	g.AddNode("A")
	g.AddNode("B")
	g.AddNode("C")
	g.AddNode("D")
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")

	fmt.Println("Graph:")
	g.PrintGraph()
	fmt.Println("BFS from A:", g.BFS("A"))
}
