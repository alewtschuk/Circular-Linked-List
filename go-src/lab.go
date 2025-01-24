package lab

import "unsafe"

// A node in the list
type Node struct {
	data unsafe.Pointer // Pointer to represent the data in the node. Equivalent to *void in C
	next *Node          // Pointer to the next node in the list
	prev *Node          // Pointer to the previous node in the list
}

// Struct to represent a list. The list maintains 2 function pointers to help
// with the management of the data it is storing. These functions must be provided by the
// user of this library.
type List struct {
	destroyData func(data unsafe.Pointer)     // Function to free the data in a node
	compareTo   func(a, b unsafe.Pointer) int // Function to compare two nodes returns 0 if the data is the same
	size        uint                          // How many nodes are in the list. Unsigned as size has to be 0 or above
	head        *Node                         // Pointer to the first(sentinal) node in the list
}

func ListInit(destroyData func(data unsafe.Pointer), compareTo func(a, b unsafe.Pointer) int) *List {
	// Initalize the head node
	var head *Node = &Node{
		data: nil,
		next: nil,
		prev: nil,
	}

	return &List{
		destroyData: destroyData,
		compareTo:   compareTo,
		size:        0,
		head:        head,
	}
}
