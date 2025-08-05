package binarySearchTree

import "fmt"

// Node represents a node in the binary search tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BinarySearchTree represents the binary search tree structure
type BinarySearchTree struct {
	Root *Node
}

// Insert adds a new value to the BST
func (t *BinarySearchTree) Insert(value int) {
	if t.Root == nil {
		t.Root = &Node{Value: value}
		return
	}
	insertNode(t.Root, value)
}

// insertNode recursively inserts a value into the BST
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

// Search looks for a value in the BST
func (t *BinarySearchTree) Search(value int) bool {
	return searchNode(t.Root, value)
}

// searchNode recursively searches for a value
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

// Delete removes a value from the BST
func (t *BinarySearchTree) Delete(value int) {
	t.Root = deleteNode(t.Root, value)
}

// deleteNode recursively deletes a node and returns the new subtree root
func deleteNode(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
		return node
	}
	if value > node.Value {
		node.Right = deleteNode(node.Right, value)
		return node
	}

	// Node to delete found
	// Case 1: No children
	if node.Left == nil && node.Right == nil {
		return nil
	}
	// Case 2: One child
	if node.Left == nil {
		return node.Right
	}
	if node.Right == nil {
		return node.Left
	}
	// Case 3: Two children
	minNode := findMin(node.Right)
	node.Value = minNode.Value
	node.Right = deleteNode(node.Right, minNode.Value)
	return node
}

// findMin finds the minimum value node in a subtree
func findMin(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// InOrderTraversal performs in-order traversal (left-root-right)
func (t *BinarySearchTree) InOrderTraversal() {
	inOrder(t.Root)
	fmt.Println()
}

// inOrder performs recursive in-order traversal
func inOrder(node *Node) {
	if node != nil {
		inOrder(node.Left)
		fmt.Printf("%d ", node.Value)
		inOrder(node.Right)
	}
}

// PreOrderTraversal performs pre-order traversal (root-left-right)
func (t *BinarySearchTree) PreOrderTraversal() {
	preOrder(t.Root)
	fmt.Println()
}

// preOrder performs recursive pre-order traversal
func preOrder(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.Value)
		preOrder(node.Left)
		preOrder(node.Right)
	}
}

// PostOrderTraversal performs post-order traversal (left-right-root)
func (t *BinarySearchTree) PostOrderTraversal() {
	postOrder(t.Root)
	fmt.Println()
}

// postOrder performs recursive post-order traversal
func postOrder(node *Node) {
	if node != nil {
		postOrder(node.Left)
		postOrder(node.Right)
		fmt.Printf("%d ", node.Value)
	}
}
