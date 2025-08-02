package doublyLinkedList

import "fmt"

type Node struct {
	data       int
	prev, next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (l *DoublyLinkedList) Insert(data int) {
	newNode := &Node{data: data, prev: nil, next: nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}
	newNode.prev = l.tail
	l.tail.next = newNode
	l.tail = newNode
}

func (l *DoublyLinkedList) Delete(data int) {
	if l.head == nil {
		return
	}
	current := l.head
	for current.data == data {
		l.head = current.next
		if l.head != nil {
			l.head.prev = nil
		} else {
			l.tail = nil
		}
		return
	}
	for current != nil && current.data != data {
		current = current.next
	}
	if current == nil {
		return
	}
	if current == l.tail {
		l.tail = current.prev
		l.tail.next = nil
		return
	}
	current.prev.next = current.next
	current.next.prev = current.prev
}

func (l *DoublyLinkedList) Display() {
	current := l.head
	if current == nil {
		fmt.Println("Doubly linked list is empty")
		return
	}
	fmt.Print("Doubly linked list: ")

	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}
