package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDA/doublyLinkedList"
)

func main() {
	list := doublyLinkedList.DoublyLinkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	fmt.Println(list.Search(3))
	list.Display()
	list.Delete(3)
	list.Display()
}
