package linkedList

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(data int) {
	newNode := &Node{data: data, next: nil}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *LinkedList) Delete(data int) {
	if l.head == nil {
		return
	}
	if l.head.data == data {
		l.head = l.head.next
		return
	}

	current := l.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func (l *LinkedList) Display() {
	current := l.head
	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}
	fmt.Print("Linked list: ")
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}
