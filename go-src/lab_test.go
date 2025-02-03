package main

import (
	"fmt"
	"testing"
	"unsafe"
)

// Utility function to allocate integer data
func allocData(i int) unsafe.Pointer {
	return unsafe.Pointer(&i)
}

// Utility function to populate the list
func populateList(lst *List) {
	for i := 0; i < 5; i++ {
		lst.ListAdd(allocData(i))
	}
}

// Test list creation and destruction
func TestCreateDestroy(t *testing.T) {
	lst := ListInit(CompareTo)
	if lst == nil {
		t.Fatalf("List should not be nil after initialization")
	}
	if lst.Head == nil {
		t.Fatalf("Sentinel node should not be nil")
	}
	if lst.Size != 0 {
		t.Fatalf("List size should be 0 after initialization")
	}
	if lst.Head.Data != nil {
		t.Fatalf("Sentinel node data should be nil")
	}

	// Test circularity
	if lst.Head.Next == nil || lst.Head.Prev == nil {
		t.Fatalf("Sentinel node's next and prev should not be nil")
	}
	if lst.Head.Next != lst.Head.Prev {
		t.Fatalf("Sentinel node's next and prev should point to each other")
	}

	// Destroy the list
	lst.ListDestroy()
	if lst.Size != 0 {
		t.Fatalf("List size should be 0 after destruction")
	}
}

// Test adding a single node
func TestAdd1(t *testing.T) {
	lst := ListInit(CompareTo)
	lst.ListAdd(allocData(1))

	if lst.Size != 1 {
		t.Fatalf("List size should be 1 after adding a node")
	}

	// Check circularity with one node
	if lst.Head.Next != lst.Head.Prev {
		t.Fatalf("With one node, head.next and head.prev should be equal")
	}
	if lst.Head == lst.Head.Next {
		t.Fatalf("Sentinel node should not be the same as the first node")
	}

	// Check data in the node
	if *(*int)(lst.Head.Next.Data) != 1 {
		t.Fatalf("Node data should be 1")
	}
}

// Test adding two nodes
func TestAdd2(t *testing.T) {
	lst := ListInit(CompareTo)
	lst.ListAdd(allocData(1))
	lst.ListAdd(allocData(2))

	if lst.Size != 2 {
		t.Fatalf("List size should be 2 after adding two nodes")
	}

	// Check circularity with two nodes
	if lst.Head.Next == lst.Head.Prev {
		t.Fatalf("With two nodes, head.next and head.prev should not be equal")
	}
	if lst.Head == lst.Head.Next {
		t.Fatalf("Sentinel node should not be the same as the first node")
	}

	// Check data in the nodes
	if *(*int)(lst.Head.Next.Data) != 2 {
		t.Fatalf("First node data should be 2")
	}
	if *(*int)(lst.Head.Prev.Data) != 1 {
		t.Fatalf("Last node data should be 1")
	}
}

// Test removing the first node
func TestRemoveIndex0(t *testing.T) {
	lst := ListInit(CompareTo)
	populateList(lst)

	removedData := lst.ListRemoveIndex(0)
	if lst.Size != 4 {
		t.Fatalf("List size should be 4 after removing one node")
	}
	if *(*int)(removedData) != 4 {
		t.Fatalf("Removed data should be 4")
	}

	// Check remaining list order: 3->2->1->0
	current := lst.Head.Next
	for i := 3; i >= 0; i-- {
		if *(*int)(current.Data) != i {
			t.Fatalf("Expected node data %d, got %d", i, *(*int)(current.Data))
		}
		current = current.Next
	}
}

// Test removing the last node
func TestRemoveIndex4(t *testing.T) {
	lst := ListInit(CompareTo)
	populateList(lst)

	removedData := lst.ListRemoveIndex(4)
	if lst.Size != 4 {
		t.Fatalf("List size should be 4 after removing one node")
	}
	if *(*int)(removedData) != 0 {
		t.Fatalf("Removed data should be 0")
	}

	// Check remaining list order: 4->3->2->1
	current := lst.Head.Next
	for i := 3; i >= 0; i-- {
		if *(*int)(current.Data) != i+1 {
			t.Fatalf("Expected node data %d, got %d", i+1, *(*int)(current.Data))
		}
		current = current.Next
	}
}

// Test invalid index removal
func TestInvalidIndex(t *testing.T) {
	lst := ListInit(CompareTo)
	populateList(lst)

	removedData := lst.ListRemoveIndex(666)
	if removedData != nil {
		t.Fatalf("Invalid index should return nil")
	}
	if lst.Size != 5 {
		t.Fatalf("List size should remain unchanged after invalid index removal")
	}
	if removedData == nil {
		fmt.Println("OUTPUT: Invalid index removal returned nil as expected")
	}
}

// Test finding an element in the list
func TestIndexOf(t *testing.T) {
	lst := ListInit(CompareTo)
	populateList(lst)

	idx := lst.ListIndexOf(allocData(3))
	if idx != 1 {
		t.Fatalf("Index of 3 should be 1, got %d", idx)
	}

	idx = lst.ListIndexOf(allocData(22))
	if idx != -1 {
		t.Fatalf("Index of non-existent element should be -1, got %d", idx)
	}
}

// Test removing all nodes
func TestRemoveAll(t *testing.T) {
	lst := ListInit(CompareTo)
	populateList(lst)

	for i := 4; i >= 0; i-- {
		removedData := lst.ListRemoveIndex(0)
		if *(*int)(removedData) != i {
			t.Fatalf("Removed data should be %d, got %d", i, *(*int)(removedData))
		}
	}

	// Ensure the list is empty
	if lst.Size != 0 {
		t.Fatalf("List size should be 0 after removing all nodes")
	}
	if lst.Head.Next != lst.Head || lst.Head.Prev != lst.Head {
		t.Fatalf("Sentinel node should point to itself after removing all nodes")
	}
}

// ADDED TESTS
func TestRemoveLastNode(t *testing.T) {
	lst := ListInit(CompareTo)
	lst.ListAdd(allocData(1)) // Add one element to the list

	removedData := lst.ListRemoveIndex(0)
	if lst.Size != 0 {
		t.Fatalf("List size should be 0 after removing the only node")
	}
	if lst.Head.Next != lst.Head || lst.Head.Prev != lst.Head {
		t.Fatalf("Sentinel node pointers should point to itself after removing the last node")
	}
	if *(*int)(removedData) != 1 {
		t.Fatalf("Removed data should be 1, got %d", *(*int)(removedData))
	}
}

func TestIndexOfMultipleOccurrences(t *testing.T) {
	lst := ListInit(CompareTo)
	lst.ListAdd(allocData(10)) // old 10
	lst.ListAdd(allocData(20))
	lst.ListAdd(allocData(10)) // new 10 at the front

	idx := lst.ListIndexOf(allocData(10))
	// The first occurrence is the brand-new 10 at index 0
	if idx != 0 {
		t.Fatalf("Index of first occurrence of 10 should be 0, got %d", idx)
	}
}

func TestEmptyList(t *testing.T) {
	lst := ListInit(CompareTo)

	// Test removal from an empty list
	removedData := lst.ListRemoveIndex(0)
	if removedData != nil {
		t.Fatalf("Removing from an empty list should return nil")
	}

	// Test index lookup in an empty list
	idx := lst.ListIndexOf(allocData(1))
	if idx != -1 {
		t.Fatalf("Index lookup in an empty list should return -1, got %d", idx)
	}

	// Ensure sentinel pointers are intact
	if lst.Head.Next != lst.Head || lst.Head.Prev != lst.Head {
		t.Fatalf("Sentinel node pointers should point to itself in an empty list")
	}

	if removedData == nil {
		fmt.Println("OUTPUT: Empty list removal returned nil as expected")
	}
}

func TestCircularTraversal(t *testing.T) {
	lst := ListInit(CompareTo)
	populateList(lst) // Populate with 4->3->2->1->0

	// Forward traversal
	current := lst.Head.Next
	for i := 4; i >= 0; i-- {
		if *(*int)(current.Data) != i {
			t.Fatalf("Forward traversal expected %d, got %d", i, *(*int)(current.Data))
		}
		current = current.Next
	}

	// Backward traversal
	current = lst.Head.Prev
	for i := 0; i <= 4; i++ {
		if *(*int)(current.Data) != i {
			t.Fatalf("Backward traversal expected %d, got %d", i, *(*int)(current.Data))
		}
		current = current.Prev
	}
}
