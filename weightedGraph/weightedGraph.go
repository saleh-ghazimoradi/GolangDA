package weightedGraph

import "fmt"

// Node represents a graph node
type Node struct {
	Value string
	Edges []*Edge // Changed from Neighbors to Edges to store weighted edges
}

// Edge represents a weighted edge to a neighbor node
type Edge struct {
	Neighbor *Node
	Weight   int // Weight of the edge
}

// Graph represents the entire weighted graph
type Graph struct {
	Nodes map[string]*Node
}

// AddNode adds a new node to the graph
func (g *Graph) AddNode(value string) {
	if _, exists := g.Nodes[value]; !exists {
		g.Nodes[value] = &Node{Value: value, Edges: []*Edge{}}
	}
}

// AddEdge adds a weighted undirected edge between two nodes
func (g *Graph) AddEdge(from, to string, weight int) {
	fromNode, fromExists := g.Nodes[from]
	toNode, toExists := g.Nodes[to]

	if !fromExists || !toExists {
		fmt.Println("One or both nodes do not exist")
		return
	}

	// Add edge from -> to if not already present
	if !containsEdge(fromNode.Edges, toNode) {
		fromNode.Edges = append(fromNode.Edges, &Edge{Neighbor: toNode, Weight: weight})
	}

	// Add edge to -> from for undirected graph
	if !containsEdge(toNode.Edges, fromNode) {
		toNode.Edges = append(toNode.Edges, &Edge{Neighbor: fromNode, Weight: weight})
	}
}

// Helper function to check if an edge to a node exists
func containsEdge(edges []*Edge, node *Node) bool {
	for _, edge := range edges {
		if edge.Neighbor == node {
			return true
		}
	}
	return false
}

// PrintGraph prints the graph's adjacency list with weights
func (g *Graph) PrintGraph() {
	for _, node := range g.Nodes {
		fmt.Printf("%s:", node.Value)
		for _, edge := range node.Edges {
			fmt.Printf(" (%s, %d)", edge.Neighbor.Value, edge.Weight)
		}
		fmt.Println()
	}
}

// NewGraph creates a new weighted graph
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}
