package treeBasedBinaryMaxHeap

// Node represents a node in the tree-based max heap
type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

// MaxHeap represents the tree-based max heap
type MaxHeap struct {
	Root *Node
	size int   // Renamed from Size to size to avoid conflict
	Last *Node // Tracks the last node for insertion
}

// NewMaxHeap creates a new empty max heap
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// Insert adds a new value to the heap
func (h *MaxHeap) Insert(value int) {
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
func (h *MaxHeap) findParentForInsert() *Node {
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

// bubbleUp moves the node up the tree to maintain the max heap property
func (h *MaxHeap) bubbleUp(node *Node) {
	if node == nil || node.Parent == nil {
		return
	}

	if node.Value > node.Parent.Value {
		node.Value, node.Parent.Value = node.Parent.Value, node.Value
		h.bubbleUp(node.Parent)
	}
}

// ExtractMax removes and returns the maximum element (root)
func (h *MaxHeap) ExtractMax() (int, bool) {
	if h.Root == nil {
		return 0, false
	}

	max := h.Root.Value
	h.size--

	if h.size == 0 {
		h.Root = nil
		h.Last = nil
		return max, true
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

	return max, true
}

// findNewLast finds the new last node after removal
func (h *MaxHeap) findNewLast() *Node {
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

// bubbleDown moves the node down the tree to maintain the max heap property
func (h *MaxHeap) bubbleDown(node *Node) {
	if node == nil {
		return
	}

	maxNode := node
	maxValue := node.Value

	if node.Left != nil && node.Left.Value > maxValue {
		maxNode = node.Left
		maxValue = node.Left.Value
	}
	if node.Right != nil && node.Right.Value > maxValue {
		maxNode = node.Right
		maxValue = node.Right.Value
	}

	if maxNode != node {
		node.Value, maxNode.Value = maxNode.Value, node.Value
		h.bubbleDown(maxNode)
	}
}

// Peek returns the maximum element without removing it
func (h *MaxHeap) Peek() (int, bool) {
	if h.Root == nil {
		return 0, false
	}
	return h.Root.Value, true
}

// Size returns the number of elements in the heap
func (h *MaxHeap) Size() int {
	return h.size
}
