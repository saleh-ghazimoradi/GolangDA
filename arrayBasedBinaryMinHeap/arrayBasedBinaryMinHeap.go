package arrayBasedBinaryMinHeap

type MinHeap struct {
	heap []int
}

// NewMinHeap creates a new min heap
func NewMinHeap() *MinHeap {
	return &MinHeap{heap: []int{}}
}

// Insert adds a new element to the heap
func (h *MinHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.bubbleUp(len(h.heap) - 1)
}

// bubbleUp moves the element at index up to its correct position
func (h *MinHeap) bubbleUp(index int) {
	parent := (index - 1) / 2
	if index <= 0 || h.heap[parent] <= h.heap[index] {
		return
	}
	h.heap[index], h.heap[parent] = h.heap[parent], h.heap[index]
	h.bubbleUp(parent)
}

// ExtractMin removes and returns the minimum element
func (h *MinHeap) ExtractMin() (int, bool) {
	if len(h.heap) == 0 {
		return 0, false
	}
	min := h.heap[0]
	last := len(h.heap) - 1
	h.heap[0] = h.heap[last]
	h.heap = h.heap[:last]
	if len(h.heap) > 0 {
		h.bubbleDown(0)
	}
	return min, true
}

// bubbleDown moves the element at index down to its correct position
func (h *MinHeap) bubbleDown(index int) {
	minIndex := index
	left := 2*index + 1
	right := 2*index + 2

	if left < len(h.heap) && h.heap[left] < h.heap[minIndex] {
		minIndex = left
	}
	if right < len(h.heap) && h.heap[right] < h.heap[minIndex] {
		minIndex = right
	}
	if minIndex != index {
		h.heap[index], h.heap[minIndex] = h.heap[minIndex], h.heap[index]
		h.bubbleDown(minIndex)
	}
}

// Peek returns the minimum element without removing it
func (h *MinHeap) Peek() (int, bool) {
	if len(h.heap) == 0 {
		return 0, false
	}
	return h.heap[0], true
}

// Size returns the number of elements in the heap
func (h *MinHeap) Size() int {
	return len(h.heap)
}
