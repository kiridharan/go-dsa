package main

import (
	"fmt"
)

type node struct {
	Data int
	next *node
}

type linkedlist struct {
	head *node
	len  int
	tail *node
}

/*
	1. Why we use * becasuse we need to work on the real data not the copy of the data
	2. make modification on the excat data . Not on Copy
	3. So that we could safe memory and also edit and view the excat linkedlist
*/
// The `prepend` function is a method of the `linkedlist` struct. It takes a pointer to a `node` as a
// parameter. This function is used to add a new node at the beginning of the linked list.
func (l *linkedlist) prepend(n *node) {
	// `temp := l.head` is creating a temporary variable `temp` and assigning it the value of `l.head`.
	// This is done to store the current head of the linked list before it is overwritten with the new
	// node `n`.

	temp := l.head
	// `l.head = n` is assigning the value of `n` to the `head` field of the `linkedlist` struct. This
	// means that the new node `n` becomes the new head of the linked list.
	l.head = n
	l.head.next = temp
	// `l.len++` is incrementing the value of the `len` field of the `linkedlist` struct by 1. This field
	// keeps track of the number of nodes in the linked list. By incrementing it, we are updating the
	// length of the linked list after adding a new node.
	l.len++

}

/*
Now here we use a copy of linkedlist because after every chnage we do her e should not affect the
real ll
*/
func (l linkedlist) printll() {
	toPrint := l.head

	for l.len != 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.next
		l.len--
	}
	fmt.Printf("\n")
}

func (l *linkedlist) deleteWithValue(value int) {

	previousToDelete := l.head

	if previousToDelete.Data == value {

		l.head.next = nil
		l.head = previousToDelete.next
	} else {

		for previousToDelete.next.Data != value {

			previousToDelete = previousToDelete.next
		}

		previousToDelete.next = previousToDelete.next.next
	}
	l.len--
}

func main() {

	mylist := linkedlist{}

	node1 := &node{Data: 344}
	mylist.tail = node1
	node2 := &node{Data: 34}
	node3 := &node{Data: 3}
	node4 := &node{Data: 0}

	mylist.prepend(node1)
	fmt.Println(mylist.len)
	mylist.prepend(node2)
	fmt.Println(mylist.len)
	mylist.prepend(node3)
	fmt.Println(mylist.len)
	mylist.prepend(node4)
	fmt.Println(mylist.len)
	fmt.Println("THE LIST")
	mylist.printll()
	fmt.Println(mylist)
	mylist.deleteWithValue(34)
	mylist.printll()
	fmt.Println(mylist)

}
