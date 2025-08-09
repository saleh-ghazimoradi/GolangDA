package graph

import "fmt"

// 1. Object-oriented based graph implementation

// Node represents a graph node
type Node struct {
	Value     string
	Neighbors []*Node
}

// Graph represents the entire graph
type Graph struct {
	Nodes map[string]*Node
}

// AddNode adds a new node to the graph
func (g *Graph) AddNode(value string) {
	if _, exists := g.Nodes[value]; !exists {
		g.Nodes[value] = &Node{Value: value}
	}
}

// AddEdge adds an undirected edge between two nodes
func (g *Graph) AddEdge(from, to string) {
	fromNode, fromExists := g.Nodes[from]
	toNode, toExists := g.Nodes[to]

	if !fromExists || !toExists {
		fmt.Println("One or both nodes do not exist")
		return
	}

	if !contains(fromNode.Neighbors, toNode) {
		fromNode.Neighbors = append(fromNode.Neighbors, toNode)
	}

	if !contains(toNode.Neighbors, fromNode) {
		toNode.Neighbors = append(toNode.Neighbors, fromNode)
	}

}

// Helper function to check if a node is in the neighbors list
func contains(neighbors []*Node, node *Node) bool {
	for _, n := range neighbors {
		if n == node {
			return true
		}
	}
	return false
}

// PrintGraph prints the graph's adjacency list
func (g *Graph) PrintGraph() {
	for _, node := range g.Nodes {
		fmt.Printf("%s:", node.Value)
		for _, neighbor := range node.Neighbors {
			fmt.Printf(" %s ", neighbor.Value)
		}
		fmt.Println()
	}
}

// NewGraph creates a new graph
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}
