package linkedlist

import (
	"strconv"
	"testing"
)

var listLength = 60000

func BenchmarkAppendBefore(b *testing.B) {
	// Create a sample linked list with a specific size
	list := &SinglyLinkedList{}
	for i := range listLength {
		list.AppendBeforeStart(strconv.Itoa(i))
	}

	node := strconv.Itoa(listLength / 2)

	// Benchmark the AppendBefore function
	b.ResetTimer()
	for range b.N {
		list.AppendBefore("B", node)
	}

}

func BenchmarkAppendAfter(b *testing.B) {
	// Create a sample linked list with a specific size
	list := &SinglyLinkedList{}
	for i := range listLength {
		list.AppendBeforeStart(strconv.Itoa(i))
	}

	node := strconv.Itoa(listLength - 1)

	// Benchmark the AppendAfter function
	b.ResetTimer()
	for range b.N {
		list.AppendAfter("A", node)
	}

}

func BenchmarkReverseListLength(b *testing.B) {
	// Create a sample linked list with a specific size
	list := &SinglyLinkedList{}

	for i := range listLength {
		list.AppendBeforeStart(strconv.Itoa(i))
	}

	// Benchmark the Reverse function
	b.ResetTimer()
	list.Reverse()

}

func BenchmarkReverse100k(b *testing.B) {
	// Create a sample linked list with a specific size
	list := &SinglyLinkedList{}
	for i := range 100000 {
		list.AppendBeforeStart(strconv.Itoa(i))
	}

	// Benchmark the Reverse function
	b.ResetTimer()
	list.Reverse()

}

func BenchmarkAddBeforeStart(b *testing.B) {
	// Create a sample linked list with a specific size
	list := &SinglyLinkedList{}
	for i := range listLength {
		list.AppendBeforeStart(strconv.Itoa(i))
	}

	// Benchmark the Reverse function
	b.ResetTimer()
	list.AppendBeforeStart("A")

}

func BenchmarkAddAfterEnd(b *testing.B) {
	// Create a sample linked list with a specific size
	list := &SinglyLinkedList{}
	for i := range listLength {
		list.AppendBeforeStart(strconv.Itoa(i))
	}

	// Benchmark the Reverse function
	b.ResetTimer()
	list.AppendToEnd("A")

}

func BenchmarkDeleteHead(b *testing.B) {
	list := &SinglyLinkedList{}

	for i := range listLength {
		list.AppendBeforeStart(strconv.Itoa(i))
	}

	b.ResetTimer()
	list.DeleteHead()

}

func BenchmarkDeleteTail(b *testing.B) {
	list := &SinglyLinkedList{}

	for i := range listLength {
		list.AppendBeforeStart(strconv.Itoa(i))
	}

	b.ResetTimer()
	list.DeleteTail()
}

func BenchmarkDeleteNode(b *testing.B) {
	list := &SinglyLinkedList{}

	for i := range listLength {
		list.AppendToEnd(strconv.Itoa(i))
	}

	nodeToDelete := strconv.Itoa(listLength / 2)

	b.ResetTimer()
	list.DeleteNode(nodeToDelete)

}

func BenchmarkDeleteOneBefore(b *testing.B) {
	list := &SinglyLinkedList{}

	for i := range listLength {
		list.AppendToEnd(strconv.Itoa(i))
	}

	nodeToDeleteBefore := strconv.Itoa(listLength / 2)

	b.ResetTimer()
	list.DeleteOneBefore(nodeToDeleteBefore)

}

func BenchmarkDeleteOneAfter(b *testing.B) {
	list := &SinglyLinkedList{}

	for i := range listLength {
		list.AppendToEnd(strconv.Itoa(i))
	}

	nodeToDeleteAfter := strconv.Itoa(listLength / 2)

	b.ResetTimer()
	list.DeleteOneAfter(nodeToDeleteAfter)
}

func BenchmarkDeleteAllBefore(b *testing.B) {
	list := &SinglyLinkedList{}

	for i := range listLength {
		list.AppendBeforeStart(strconv.Itoa(i))
	}

	nodeToDeleteAllBefore := strconv.Itoa(listLength / 2)

	b.ResetTimer()
	list.DeleteAllBefore(nodeToDeleteAllBefore)

}

func BenchmarkDeleteAllAfter(b *testing.B) {
	list := &SinglyLinkedList{}

	for i := range listLength {
		list.AppendBeforeStart(strconv.Itoa(i))
	}

	nodeToDeleteAllAfter := strconv.Itoa(listLength / 2)

	b.ResetTimer()
	list.DeleteAllAfter(nodeToDeleteAllAfter)
}
