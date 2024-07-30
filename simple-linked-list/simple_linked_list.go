package linkedlist

import "errors"

// Define the List and Element types here.

// Node represents a node in the doubly linked list.
type Node struct {
	Value int
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
func New(elements []int) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	newNode := &Node{Value: element}
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

func (l *List) Pop() (int, error) {
	if l.tail == nil {
		return 0, errors.New("list is empty")
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

func (l *List) Array() []int {
	var result []int
	for n := l.head; n != nil; n = n.next {
		result = append(result, n.Value)
	}
	return result
}

func (l *List) Reverse() *List {
	newList := &List{}
	for n := l.tail; n != nil; n = n.prev {
		newList.Push(n.Value)
	}
	return newList
}
