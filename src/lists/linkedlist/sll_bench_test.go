package linkedlist

import (
	"strconv"
	"testing"
)

var listLength = 60000

func populateList(len int) *SinglyLinkedList {

	list := &SinglyLinkedList{}

	for i := range len {
		list.AppendBeforeStart(strconv.Itoa(i))
	}

	return list

}

func BenchmarkAppendBefore(b *testing.B) {
	// Create a sample linked list with a specific size
	list := populateList(listLength)

	node := strconv.Itoa(listLength / 2)

	// Benchmark the AppendBefore function
	b.ResetTimer()
	for range b.N {
		list.AppendBefore("B", node)
	}

}

func BenchmarkAppendAfter(b *testing.B) {

	list := populateList(listLength)

	node := strconv.Itoa(listLength - 1)

	b.ResetTimer()

	for range b.N {
		list.AppendAfter("A", node)
	}

}

func BenchmarkReverse(b *testing.B) {

	b.Run("Reverse100k", func(b *testing.B) {

		list := &SinglyLinkedList{}

		for i := range 100000 {
			list.AppendBeforeStart(strconv.Itoa(i))
		}

		b.ResetTimer()

		list.Reverse()

	})

	b.Run("Reverse10k", func(b *testing.B) {

		list := &SinglyLinkedList{}

		for i := range 10000 {
			list.AppendBeforeStart(strconv.Itoa(i))
		}

		b.ResetTimer()

		list.Reverse()

	})

}

func BenchmarkAddBeforeStart(b *testing.B) {

	list := populateList(listLength)

	b.ResetTimer()

	list.AppendBeforeStart("A")

}

func BenchmarkAddAfterEnd(b *testing.B) {

	list := populateList(listLength)

	b.ResetTimer()

	list.AppendToEnd("A")

}

func BenchmarkDeleteHead(b *testing.B) {

	list := populateList(listLength)

	b.ResetTimer()

	list.DeleteHead()

}

func BenchmarkDeleteTail(b *testing.B) {

	list := populateList(listLength)

	b.ResetTimer()

	list.DeleteTail()

}

func BenchmarkDeleteNode(b *testing.B) {

	list := populateList(listLength)

	b.Run("HeadNode", func(b *testing.B) {

		b.ResetTimer()

		list.DeleteNode("0")

	})

	b.Run("MiddleNode", func(b *testing.B) {

		node := strconv.Itoa(listLength / 2)

		b.ResetTimer()

		list.DeleteNode(node)

	})

	b.Run("TailNode", func(b *testing.B) {

		node := strconv.Itoa(listLength - 3)

		b.ResetTimer()

		list.DeleteNode(node)

	})

	b.Run("NonExistentNode", func(b *testing.B) {

		node := strconv.Itoa(listLength + 3)

		b.ResetTimer()

		list.DeleteNode(node)

	})

}

func BenchmarkDeleteOneBefore(b *testing.B) {

	b.Run("HeadNode", func(b *testing.B) {

		list := populateList(listLength)

		nodeToDeleteBefore := strconv.Itoa(listLength / 2)

		b.ResetTimer()

		list.DeleteOneBefore(nodeToDeleteBefore)

	})

	b.Run("MiddleNode", func(b *testing.B) {

		list := populateList(listLength)

		node := strconv.Itoa(listLength / 2)

		b.ResetTimer()

		list.DeleteOneBefore(node)

	})

	b.Run("TailNode", func(b *testing.B) {

		list := populateList(listLength)

		node := strconv.Itoa(listLength - 1)

		b.ResetTimer()

		list.DeleteOneBefore(node)

	})

}

func BenchmarkDeleteOneAfter(b *testing.B) {

	list := populateList(listLength)

	nodeToDeleteAfter := strconv.Itoa(listLength / 2)

	b.ResetTimer()

	list.DeleteOneAfter(nodeToDeleteAfter)

}

func BenchmarkDeleteAllBefore(b *testing.B) {

	list := populateList(listLength)

	nodeToDeleteAllBefore := strconv.Itoa(listLength / 2)

	b.ResetTimer()

	list.DeleteAllBefore(nodeToDeleteAllBefore)

}

func BenchmarkDeleteAllAfter(b *testing.B) {

	list := populateList(listLength)

	nodeToDeleteAllAfter := strconv.Itoa(listLength / 2)

	b.ResetTimer()

	list.DeleteAllAfter(nodeToDeleteAllAfter)

}
