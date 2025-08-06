package treeBasedBinaryMinHeap

// Node represents a node in the tree-based min heap
type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

// MinHeap represents the tree-based min heap
type MinHeap struct {
	Root *Node
	size int
	Last *Node // Tracks the last node for insertion
}

// NewMinHeap creates a new empty min heap
func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

// Insert adds a new value to the heap
func (h *MinHeap) Insert(value int) {
	newNode := &Node{Value: value}
	h.size++

	if h.Root == nil {
		h.Root = newNode
		h.Last = newNode
		return
	}

	// Find the parent for the new node (next available spot in level-order)
	parent := h.findParentForInsert()
	newNode.Parent = parent

	// Attach the new node to the parent
	if parent.Left == nil {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}
	h.Last = newNode

	// Bubble up to maintain heap property
	h.bubbleUp(newNode)
}

// findParentForInsert finds the parent node for the next insertion
func (h *MinHeap) findParentForInsert() *Node {
	if h.Last == nil {
		return nil
	}

	// Use level-order traversal to find the first parent with an empty child slot
	queue := []*Node{h.Root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Left == nil || current.Right == nil {
			return current
		}

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return nil
}

// bubbleUp moves the node up the tree to maintain the min heap property
func (h *MinHeap) bubbleUp(node *Node) {
	if node == nil || node.Parent == nil {
		return
	}

	if node.Value < node.Parent.Value {
		node.Value, node.Parent.Value = node.Parent.Value, node.Value
		h.bubbleUp(node.Parent)
	}
}

// ExtractMin removes and returns the minimum element (root)
func (h *MinHeap) ExtractMin() (int, bool) {
	if h.Root == nil {
		return 0, false
	}

	min := h.Root.Value
	h.size--

	if h.size == 0 {
		h.Root = nil
		h.Last = nil
		return min, true
	}

	// Replace root with the last node's value
	h.Root.Value = h.Last.Value

	// Remove the last node
	parent := h.Last.Parent
	if parent != nil {
		if parent.Left == h.Last {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	}

	// Update the last node
	h.Last = h.findNewLast()

	// Bubble down from the root
	h.bubbleDown(h.Root)

	return min, true
}

// findNewLast finds the new last node after removal
func (h *MinHeap) findNewLast() *Node {
	if h.size == 0 {
		return nil
	}

	// Use level-order traversal to find the last node
	queue := []*Node{h.Root}
	var last *Node
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		last = current

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return last
}

// bubbleDown moves the node down the tree to maintain the min heap property
func (h *MinHeap) bubbleDown(node *Node) {
	if node == nil {
		return
	}

	minNode := node
	minValue := node.Value

	if node.Left != nil && node.Left.Value < minValue {
		minNode = node.Left
		minValue = node.Left.Value
	}
	if node.Right != nil && node.Right.Value < minValue {
		minNode = node.Right
		minValue = node.Right.Value
	}

	if minNode != node {
		node.Value, minNode.Value = minNode.Value, node.Value
		h.bubbleDown(minNode)
	}
}

// Peek returns the minimum element without removing it
func (h *MinHeap) Peek() (int, bool) {
	if h.Root == nil {
		return 0, false
	}
	return h.Root.Value, true
}

// Size returns the number of elements in the heap
func (h *MinHeap) Size() int {
	return h.size
}
