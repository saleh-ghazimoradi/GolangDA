package binaryTree

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (b *BinaryTree) Insert(value int) {
	if b.Root == nil {
		b.Root = &Node{Value: value}
		return
	}
	insertNode(b.Root, value)
}

func insertNode(node *Node, value int) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
		} else {
			insertNode(node.Left, value)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value}
		} else {
			insertNode(node.Right, value)
		}
	}
}

func (b *BinaryTree) InOrderTraversal() {
	inOrder(b.Root)
	fmt.Println()
}

func inOrder(node *Node) {
	if node != nil {
		inOrder(node.Left)
		fmt.Printf("%d ", node.Value)
		inOrder(node.Right)
	}
}

func (b *BinaryTree) Search(value int) bool {
	return searchNode(b.Root, value)
}

func searchNode(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		return true
	}
	if value < node.Value {
		return searchNode(node.Left, value)
	}
	return searchNode(node.Right, value)
}
