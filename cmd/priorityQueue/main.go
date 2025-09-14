package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/priorityQueue"
)

func main() {
	pq := priorityQueue.NewPriorityQueue()
	pq.Push(3)
	pq.Push(1)
	pq.Push(4)
	pq.Push(2)

	fmt.Println("Priority Queue Size:", pq.Size())
	if val, ok := pq.Peek(); ok {
		fmt.Println("Top element:", val)
	}

	fmt.Println("Popped elements:")
	for pq.Size() > 0 {
		if val, ok := pq.Pop(); ok {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()
}
