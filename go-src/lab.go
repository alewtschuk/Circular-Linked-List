// Author: Alex Lewtschuk
// Program: A cirtular doubly linked list in Go, translated from C
package main

import (
	"fmt"
	"unsafe"
)

// A node in the list
type Node struct {
	Data unsafe.Pointer // Pointer to represent the data in the node. Equivalent to *void in C
	Next *Node          // Pointer to the Next node in the list
	Prev *Node          // Pointer to the Previous node in the list
}

// Struct to represent a list. The list maintains 2 function pointers to help
// with the management of the data it is storing. These functions must be provided by the
// user of this library.
type List struct {
	compareTo func(a, b unsafe.Pointer) int // Function to compare two nodes returns 0 if the data is the same
	Size      uint64                        // How many nodes are in the list. Unsigned as Size has to be 0 or above
	Head      *Node                         // Pointer to the first(sentinal) node in the list
}

// ListInit initializes the list with the provided function to compare the data
func ListInit(compareTo func(a, b unsafe.Pointer) int) *List {
	// Initalize the sentinal node
	var sentinal *Node = &Node{
		Data: nil,
		Next: nil,
		Prev: nil,
	}

	sentinal.Next = sentinal // Set the Next node of the sentinal node to itself
	sentinal.Prev = sentinal // Set the Previous node of the sentinal node to itself

	return &List{
		compareTo: compareTo,
		Size:      0,
		Head:      sentinal, // Set the Head of the list to the sentinal node
	}
}

// ListDestroy destroys the list and frees all the memory
func (l *List) ListDestroy() {
	// Start from the first node in the list
	var current *Node = l.Head.Next

	// Iterate through all nodes in the list
	for l.Size > 0 {
		// Store the Next node
		var Next *Node = current.Next

		// Break the links of the current node
		current.Data = nil
		current.Next = nil
		current.Prev = nil

		// Move to the Next node
		current = Next

		// Decrement the Size of the list
		l.Size--
	}

	// Reset the sentinel node links
	l.Head.Next = l.Head
	l.Head.Prev = l.Head
}

// ListAdd adds a new node to the front of the list
func (l *List) ListAdd(data unsafe.Pointer) {
	// Create new node to be added to the front of the list
	var newNode *Node = &Node{
		Data: data,        // Set the data of the new node to the data provided
		Next: l.Head.Next, // Set the Next node of the new node to the current first node
		Prev: l.Head,      // Set the Previous node of the new node to the sentinal node
	}

	if l.Size == 0 { // If the list is empty establish circular links
		newNode.Next = l.Head // Set the Next node of the first new node added to the list to the sentinal node
		newNode.Prev = l.Head // Set the Previous node of the first new node added to the list to the sentinal node

		l.Head.Next = newNode // Set the Next node of the sentinal node to the new node
		l.Head.Prev = newNode // Set the Previous node of the sentinal node to the new node
	} else { // The list is not empty
		// l.Head.Next.Prev is the Previous node of the current first node
		l.Head.Next.Prev = newNode // Set the Previous node of the current first node to the new node
		newNode.Prev = l.Head      // Set the Previous node of the new node to the sentinal node
		newNode.Next = l.Head.Next // Set the Next node of the new node to the current first node
		l.Head.Next = newNode      // Set the Next node of the sentinal node to the new node
	}
	l.Size++ // Increment the Size of the list
}

// ListRemoveIndex removes the node at the index and returns the data
func (l *List) ListRemoveIndex(index uint64) unsafe.Pointer {
	// Check if the list is empty or the index is out of bounds
	if l.Size == 0 {
		fmt.Println("Cannot remove from an empty list")
		return nil
	}
	if index >= l.Size {
		fmt.Println("Index out of bounds")
		return nil
	}

	var current *Node = l.Head.Next // Start at the first node in the list
	for i := uint64(0); i < index; i++ {
		current = current.Next // Itterate through the list to the node at the index
	}

	// If the list only has one node
	if l.Size == 1 {
		l.Head.Next = l.Head
		l.Head.Prev = l.Head
	} else {
		//Update the links of the nodes around the node being removed
		current.Prev.Next = current.Next // Set the Next node of the Previous node to the node after the node being removed
		current.Next.Prev = current.Prev // Set the Previous node of the Next node to the node before the node being removed
	}
	l.Size-- // Decrement the Size of the list
	return current.Data
}

// ListIndexOf returns the index of the first occurrence of the data in the list
func (l *List) ListIndexOf(data unsafe.Pointer) int {
	if l.Size == 0 {
		return -1
	}

	var current *Node = l.Head.Next // Start at the first node in the list
	for index := uint64(0); index < l.Size; index++ {
		if l.compareTo(current.Data, data) == 0 {
			return int(index)
		}
		current = current.Next
	}
	return -1 // Data not found
}

// CompareTo compares two integers and returns 0 if they are equal
func CompareTo(a, b unsafe.Pointer) int {
	var valA int = *(*int)(a)
	var valB int = *(*int)(b)

	if valA == valB {
		return 0
	} else {
		return 1
	}
}
