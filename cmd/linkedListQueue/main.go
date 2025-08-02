package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/linkedListQueue"
)

func main() {
	queue := linkedListQueue.Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Display()
	if data, ok := queue.Dequeue(); ok {
		fmt.Printf("Dequeued: %d\n", data)
	}
	queue.Display()
}
