package dijkstra

import (
	"math"
)

// Edge represents a weighted edge in the graph
type Edge struct {
	To     int
	Weight int
}

// Graph represents an adjacency list graph
type Graph struct {
	Vertices int
	AdjList  map[int][]Edge
}

// NewGraph creates a new graph with the given number of Vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		AdjList:  make(map[int][]Edge),
	}
}

// AddEdge adds a weighted edge To the graph
func (g *Graph) AddEdge(from, To, weight int) {
	g.AdjList[from] = append(g.AdjList[from], Edge{To, weight})
	// For undirected graph, add the reverse edge
	g.AdjList[To] = append(g.AdjList[To], Edge{from, weight})
}

// Dijkstra implements Dijkstra's algorithm To find shortest paths from a source vertex
func (g *Graph) Dijkstra(source int) ([]int, []int) {
	// Initialize distances and predecessors
	dist := make([]int, g.Vertices)
	prev := make([]int, g.Vertices)
	visited := make([]bool, g.Vertices)

	for i := 0; i < g.Vertices; i++ {
		dist[i] = math.MaxInt32
		prev[i] = -1
	}
	dist[source] = 0

	// Priority queue simulation using a min-heap
	type Node struct {
		Vertex int
		Dist   int
	}

	pq := make([]Node, 0)
	pq = append(pq, Node{source, 0})

	for len(pq) > 0 {
		// Extract minimum distance vertex
		minNode := pq[0]
		pq = pq[1:]

		curr := minNode.Vertex
		if visited[curr] {
			continue
		}

		visited[curr] = true

		// Process all neighbors
		for _, edge := range g.AdjList[curr] {
			if !visited[edge.To] && dist[curr]+edge.Weight < dist[edge.To] {
				dist[edge.To] = dist[curr] + edge.Weight
				prev[edge.To] = curr
				pq = append(pq, Node{edge.To, dist[edge.To]})

				// Sort pq by distance (simple bubble sort for small graphs)
				for i := len(pq) - 1; i > 0; i-- {
					if pq[i].Dist < pq[i-1].Dist {
						pq[i], pq[i-1] = pq[i-1], pq[i]
					}
				}
			}
		}
	}

	return dist, prev
}
