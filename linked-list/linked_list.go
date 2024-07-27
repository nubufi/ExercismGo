package linkedlist

import "errors"

// Node represents a node in the doubly linked list.
type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}

// Next returns the next node.
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns the previous node.
func (n *Node) Prev() *Node {
	return n.prev
}

// List represents the doubly linked list.
type List struct {
	head *Node
	tail *Node
	size int
}

// NewList creates a new linked list preserving the order of the values.
func NewList(elements ...interface{}) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

// First returns a pointer to the first node (head).
func (l *List) First() *Node {
	return l.head
}

// Last returns a pointer to the last node (tail).
func (l *List) Last() *Node {
	return l.tail
}

// Push inserts value at the back of the list.
func (l *List) Push(v interface{}) {
	newNode := &Node{Value: v}
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

// Pop removes value from the back of the list.
func (l *List) Pop() (interface{}, error) {
	if l.tail == nil {
		return nil, errors.New("list is empty")
	}
	value := l.tail.Value
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}
	l.size--
	return value, nil
}

// Unshift inserts value at the front of the list.
func (l *List) Unshift(v interface{}) {
	newNode := &Node{Value: v}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
	l.size++
}

// Shift removes value from the front of the list.
func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("list is empty")
	}
	value := l.head.Value
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}
	l.size--
	return value, nil
}

// Reverse reverses the linked list.
func (l *List) Reverse() {
	if l.head == nil || l.head == l.tail {
		return
	}
	current := l.head
	var prev *Node
	l.tail = l.head
	for current != nil {
		next := current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}
	l.head = prev
}

