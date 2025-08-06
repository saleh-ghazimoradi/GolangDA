package arrayBasedBinaryMaxHeap

type MaxHeap struct {
	heap []int
}

// NewMaxHeap creates a new max heap
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{heap: []int{}}
}

// Insert adds a new element to the heap
func (h *MaxHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.bubbleUp(len(h.heap) - 1)
}

// bubbleUp moves the element at index up to its correct position
func (h *MaxHeap) bubbleUp(index int) {
	parent := (index - 1) / 2
	if index <= 0 || h.heap[parent] >= h.heap[index] {
		return
	}
	h.heap[index], h.heap[parent] = h.heap[parent], h.heap[index]
	h.bubbleUp(parent)
}

// ExtractMax removes and returns the maximum element
func (h *MaxHeap) ExtractMax() (int, bool) {
	if len(h.heap) == 0 {
		return 0, false
	}
	max := h.heap[0]
	last := len(h.heap) - 1
	h.heap[0] = h.heap[last]
	h.heap = h.heap[:last]
	if len(h.heap) > 0 {
		h.bubbleDown(0)
	}
	return max, true
}

// bubbleDown moves the element at index down to its correct position
func (h *MaxHeap) bubbleDown(index int) {
	maxIndex := index
	left := 2*index + 1
	right := 2*index + 2

	if left < len(h.heap) && h.heap[left] > h.heap[maxIndex] {
		maxIndex = left
	}
	if right < len(h.heap) && h.heap[right] > h.heap[maxIndex] {
		maxIndex = right
	}
	if maxIndex != index {
		h.heap[index], h.heap[maxIndex] = h.heap[maxIndex], h.heap[index]
		h.bubbleDown(maxIndex)
	}
}

// Peek returns the maximum element without removing it
func (h *MaxHeap) Peek() (int, bool) {
	if len(h.heap) == 0 {
		return 0, false
	}
	return h.heap[0], true
}

// Size returns the number of elements in the heap
func (h *MaxHeap) Size() int {
	return len(h.heap)
}
