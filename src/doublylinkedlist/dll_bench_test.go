package doublylinkedlist

import (
	"strconv"
	"testing"
)

var listLength = 60000

func populateList() *DoublyLinkedList {

	list := &DoublyLinkedList{}

	for i := range listLength {
		list.Append(strconv.Itoa(i))
	}

	return list

}

func BenchmarkAppend(b *testing.B) {

	list := populateList()

	b.ResetTimer()
	for i := range b.N {
		list.Append(strconv.Itoa(i))
	}

}

func BenchmarkPrepend(b *testing.B) {

	list := populateList()

	b.ResetTimer()
	for i := range b.N {
		list.Prepend(strconv.Itoa(i))
	}

}

func BenchmarkFindNode(b *testing.B) {

	list := populateList()

	b.ResetTimer()

	b.Run("EmptyList", func(b *testing.B) {

		list = &DoublyLinkedList{}

		for range b.N {
			list.findNode("0")
		}

	})

	b.Run("HeadMatch", func(b *testing.B) {

		for range b.N {
			list.findNode("0")
		}

	})

	b.Run("MiddleMatch", func(b *testing.B) {

		node := strconv.Itoa(listLength / 2)

		b.ResetTimer()

		for range b.N {
			list.findNode(node)
		}

	})

	b.Run("TailMatch", func(b *testing.B) {

		node := strconv.Itoa(listLength - 1)

		b.ResetTimer()

		for range b.N {
			list.findNode(node)
		}

	})

	b.Run("NoMatch", func(b *testing.B) {

		node := strconv.Itoa(listLength + 1)

		b.ResetTimer()

		for range b.N {
			list.findNode(node)
		}

	})

}

func BenchmarkAppendAfter(b *testing.B) {

	list := populateList()

	b.ResetTimer()

	b.Run("EmptyList", func(b *testing.B) {

		list = &DoublyLinkedList{}

		for range b.N {
			list.AppendAfter("0", "0")
		}

	})

	b.Run("HeadMatch", func(b *testing.B) {

		for range b.N {
			list.AppendAfter("0", "0")
		}

	})

	b.Run("MiddleMatch", func(b *testing.B) {

		node := strconv.Itoa(listLength / 2)

		b.ResetTimer()

		for range b.N {
			list.AppendAfter("0", node)
		}

	})

	b.Run("TailMatch", func(b *testing.B) {

		node := strconv.Itoa(listLength - 1)

		b.ResetTimer()

		for range b.N {
			list.AppendAfter("0", node)
		}

	})

}

func BenchmarkAppendBefore(b *testing.B) {

	list := populateList()

	b.ResetTimer()

	b.Run("EmptyList", func(b *testing.B) {

		list = &DoublyLinkedList{}

		for range b.N {
			list.AppendBefore("0", "0")
		}

	})

	b.Run("HeadMatch", func(b *testing.B) {

		for range b.N {
			list.AppendBefore("0", "0")
		}

	})

	b.Run("MiddleMatch", func(b *testing.B) {

		node := strconv.Itoa(listLength / 2)

		b.ResetTimer()

		for range b.N {
			list.AppendBefore("0", node)
		}

	})

	b.Run("TailMatch", func(b *testing.B) {

		node := strconv.Itoa(listLength - 1)

		b.ResetTimer()

		for range b.N {
			list.AppendBefore("0", node)
		}

	})

}
