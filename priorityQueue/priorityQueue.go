package priorityQueue

// PriorityQueue implements a min-heap
type PriorityQueue struct {
	heap []int
}

// NewPriorityQueue creates a new priority queue
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{heap: []int{}}
}

// Push adds an element to the priority queue
func (pq *PriorityQueue) Push(val int) {
	pq.heap = append(pq.heap, val)
	pq.heapifyUp(len(pq.heap) - 1)
}

// Pop removes and returns the minimum element
func (pq *PriorityQueue) Pop() (int, bool) {
	if len(pq.heap) == 0 {
		return 0, false
	}
	min := pq.heap[0]
	last := len(pq.heap) - 1
	pq.heap[0] = pq.heap[last]
	pq.heap = pq.heap[:last]
	if len(pq.heap) > 0 {
		pq.heapifyDown(0)
	}
	return min, true
}

// Peek returns the minimum element without removing it
func (pq *PriorityQueue) Peek() (int, bool) {
	if len(pq.heap) == 0 {
		return 0, false
	}
	return pq.heap[0], true
}

// Size returns the number of elements in the priority queue
func (pq *PriorityQueue) Size() int {
	return len(pq.heap)
}

// heapifyUp maintains the heap property by moving an element up
func (pq *PriorityQueue) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if pq.heap[index] >= pq.heap[parent] {
			break
		}
		pq.heap[index], pq.heap[parent] = pq.heap[parent], pq.heap[index]
		index = parent
	}
}

// heapifyDown maintains the heap property by moving an element down
func (pq *PriorityQueue) heapifyDown(index int) {
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < len(pq.heap) && pq.heap[left] < pq.heap[smallest] {
			smallest = left
		}
		if right < len(pq.heap) && pq.heap[right] < pq.heap[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}
		pq.heap[index], pq.heap[smallest] = pq.heap[smallest], pq.heap[index]
		index = smallest
	}
}
