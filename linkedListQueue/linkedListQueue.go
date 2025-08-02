package linkedListQueue

import "fmt"

type Node struct {
	data       int
	prev, next *Node
}

type Queue struct {
	head, tail *Node
	size       int
}

func (q *Queue) Enqueue(data int) {
	newNode := &Node{data: data, prev: nil, next: nil}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		newNode.prev = q.tail
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
}

func (q *Queue) Dequeue() (int, bool) {
	if q.head == nil {
		return 0, false
	}
	data := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	} else {
		q.head.prev = nil
	}
	q.size--
	return data, true
}

func (q *Queue) Display() {
	if q.head == nil {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Print("Queue: ")
	current := q.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}
