package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/dijkstra"
)

// Example usage
func main() {

	// Create a graph with 5 vertices
	g := dijkstra.NewGraph(5)

	// Add edges
	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 2, 8)
	g.AddEdge(1, 2, 2)
	g.AddEdge(1, 3, 5)
	g.AddEdge(2, 3, 5)
	g.AddEdge(2, 4, 9)
	g.AddEdge(3, 4, 4)

	// Run Dijkstra's algorithm from vertex 0
	dist, prev := g.Dijkstra(0)

	// Print results
	fmt.Println("Shortest distances from vertex 0:")
	for i, d := range dist {
		fmt.Printf("To vertex %d: %d\n", i, d)
	}

	fmt.Println("\nPredecessor array:")
	for i, p := range prev {
		fmt.Printf("Vertex %d: %d\n", i, p)
	}
}
