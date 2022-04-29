package main

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(value interface{}) {
	n := &Node{
		next:  nil,
		value: value,
	}
	if l.head == nil {
		l.head = n
		return
	}
	node := l.head
	for node.next != nil {
		node = node.next
	}
	node.next = n
}

func (l LinkedList) Print() {
	if l.head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	fmt.Print(l.head.value)
	next := l.head.next
	for next != nil {
		fmt.Printf(" -> %v", next.value)
		next = next.next
	}
	fmt.Println()
}

func (l LinkedList) Reverse() LinkedList {
	if l.head == nil {
		return LinkedList{}
	}

	next := l.head

	n := &Node{
		next:  nil,
		value: l.head.value,
	}

	for next.next != nil {
		next = next.next
		n = &Node{
			next:  n,
			value: next.value,
		}
	}

	return LinkedList{
		head: n,
	}

}
