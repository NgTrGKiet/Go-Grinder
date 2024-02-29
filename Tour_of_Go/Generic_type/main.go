package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// PushFront inserts a new element with value 'val' at the front of the list.
func (l *List[T]) PushFront(val T) *List[T] {
	newNode := &List[T]{val: val, next: l}
	return newNode
}

// PrintList traverses the list and prints each element.
func (l *List[T]) PrintList() {
	curr := l
	for curr != nil {
		fmt.Println(curr.val)
		curr = curr.next
	}
}

func main() {
	// Create a new list
	var list *List[int]

	// Insert elements at the front of the list
	list = list.PushFront(3)
	list = list.PushFront(2)
	list = list.PushFront(1)

	// Print the list
	list.PrintList()
}
