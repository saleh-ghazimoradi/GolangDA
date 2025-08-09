package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/graphDFS"
)

func main() {
	g := graphDFS.NewGraph()
	g.AddNode("A")
	g.AddNode("B")
	g.AddNode("C")
	g.AddNode("D")
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")

	fmt.Println("Graph:")
	g.PrintGraph()

	fmt.Println("DFS from A:", g.DFS("A")) // Example: [A B D C]
}
