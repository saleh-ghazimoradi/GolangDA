package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/queue"
)

func main() {
	q := queue.NewQueue[int]()

	// Demonstrate queue operations
	fmt.Println("Queue Operations:")
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Printf("Queue size: %d\n", q.Size())

	if item, ok := q.Peek(); ok {
		fmt.Printf("Queue peek: %d\n", item)
	}

	for !q.IsEmpty() {
		if item, ok := q.Dequeue(); ok {
			fmt.Printf("Dequeued: %d\n", item)
		}
	}
	fmt.Printf("Queue is empty: %v\n", q.IsEmpty())
}
