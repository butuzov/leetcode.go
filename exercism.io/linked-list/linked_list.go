package linkedlist

// package main

// This package provide simple implementation of double linked list.
// Supported methods
// Insert
// PushBack/PushFront
// PopFront
// PopBack

import (
	"bytes"
	"errors"
	"fmt"
)

// ErrEmptyList - Error Code we use.
var ErrEmptyList = errors.New("404: Not Enough Elements to Perform Opperation")

// List - Linked List - has tail and head.
type List struct {
	head *Node
	tail *Node
}

// Node is linked list node struct
type Node struct {
	prev *Node
	next *Node
	Val  interface{}
}

// NewNode - create new Node struct with a provided value.
func NewNode(val interface{}) *Node {
	node := new(Node)
	node.Val = val
	return node
}

// String formatter for easier debug.
func (dll List) String() string {
	node := dll.head

	buf := bytes.NewBuffer([]byte{'{'})

	if node != nil {

		if dll.head == dll.tail {
			buf.WriteString(fmt.Sprintf("%+v", node.Val))
		} else {

			var n int
			var format string

			for node != nil {

				if node.next == nil {
					format = "%+v"
				} else {
					format = "%+v<->"
				}

				buf.WriteString(fmt.Sprintf(format, node.Val))
				node = node.next
			}
		}
	}

	buf.WriteByte('}')

	return buf.String()
}

// Display use string formating under the hood.
func (dll List) Display() {
	fmt.Println(dll)
}

// Next returns pointer to next node.
func (n Node) Next() *Node {
	return n.next
}

// Prev returns pointer to prev node.
func (n Node) Prev() *Node {
	return n.prev
}

// First returns pointer to first/head node.
func (dll List) First() *Node {
	return dll.head
}

// Last return pointer to last/tail node.
func (dll List) Last() *Node {
	return dll.tail
}

// Reverse will reverse order in the linked list.
func (dll *List) Reverse() {
	// var head, node *Node

	var node, next *Node

	// cur := ll.Last()
	for node = dll.Last(); node != nil; node = next {

		// preserving previous
		next = node.Prev()

		// switching next and prev pointers for node
		node.next, node.prev = node.prev, node.next
	}

	dll.head, dll.tail = dll.Last(), dll.First()

}

// Insert arbitrary value into the linked list.
func (dll *List) Insert(n *Node) {

	if dll.head == nil {

		dll.head = n
		dll.tail = n
		dll.tail.prev, dll.head.next = nil, nil

	} else {

		if dll.head == dll.tail {
			dll.head.next = n
			dll.head.prev = nil

		}

		dll.tail.next = n
		n.prev = dll.tail
		dll.tail = n
	}

}

// NewList - Returns pointer to new List
func NewList(elements ...interface{}) *List {
	list := new(List)
	for _, value := range elements {
		list.Insert(NewNode(value))
	}
	return list
}

// IsEmpty return boolean status is this empty or not.
func (dll *List) IsEmpty() bool {
	if dll.head == nil && dll.tail == nil {
		return true
	}
	return false
}

// PushBack - push arbitrary value to the end/tail of linked list.
func (dll *List) PushBack(val interface{}) {
	node := NewNode(val)
	if dll.IsEmpty() {
		dll.Insert(node)
		return
	}

	dll.Last().next = node
	node.prev = dll.Last()
	dll.tail = node
}

// PushFront - push arbitrary value to the begining/head of linked list.
func (dll *List) PushFront(val interface{}) {
	node := NewNode(val)
	if dll.IsEmpty() {
		dll.Insert(node)
		return
	}

	dll.First().prev = node
	node.next = dll.First()
	dll.head = node
}

// PopFront - pop element from the begining of list.
func (dll *List) PopFront() (interface{}, error) {
	if dll.IsEmpty() {
		return 0, ErrEmptyList
	}

	head := dll.First()

	if head.next != nil {
		dll.head = head.next
		head.next = nil
		dll.head.prev = nil
	} else {
		dll.head, dll.tail = nil, nil
	}

	return head.Val, nil
}

// PopBack - pop element from the end of list.
func (dll *List) PopBack() (interface{}, error) {
	if dll.IsEmpty() {
		return 0, ErrEmptyList
	}

	tail := dll.Last()

	if tail.prev != nil {
		dll.tail = tail.prev
		tail.prev = nil
		dll.tail.next = nil
	} else {
		dll.head, dll.tail = nil, nil
	}

	return tail.Val, nil
}

// this was tests, when it was main
// func main() {
// 	nums := []interface{}{1, 10, 100, 431}
// 	list := NewList(nums...)
// 	fmt.Printf("List: %s\n", list)
// 	list.Reverse()
// 	fmt.Printf("List Reverse: %s\n", list)

// 	list.PushBack(1)
// 	list.PushBack(2)
// 	list.PushBack(25)
// 	list.PushFront(25)
// 	fmt.Printf("Pushes: %s\n", list)

// 	nl1, error1 := list.PopFront()
// 	fmt.Printf("%d, %v - %s\n", nl1, error1, list)

// 	nl2, error2 := list.PopBack()
// 	fmt.Printf("%d, %v - %s\n", nl2, error2, list)
// }
