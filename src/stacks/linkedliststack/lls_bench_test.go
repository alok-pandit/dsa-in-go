package linkedliststack

import (
	"testing"
)

func BenchmarkPush(b *testing.B) {

	b.Run("PushToEmptyStack", func(b *testing.B) {

		stack := &LinkedListStack{}

		b.ResetTimer()

		for i := range b.N {
			stack.Push(i)
		}

	})

	b.Run("PushAtop2kFilled", func(b *testing.B) {

		stack := &LinkedListStack{}

		for i := range 2000 {
			stack.Push(i)
		}

		b.ResetTimer()

		for i := range b.N {
			stack.Push(i)
		}

	})

	b.Run("PushTo2kCapEmpty", func(b *testing.B) {

		stack := &LinkedListStack{}

		b.ResetTimer()

		for i := range b.N {
			stack.Push(i)
		}

	})

}

func BenchmarkPop(b *testing.B) {

	b.Run("PopFromEmptyStack", func(b *testing.B) {

		stack := &LinkedListStack{}

		b.ResetTimer()

		for range b.N {
			stack.Pop()
		}

	})

	b.Run("PopFrom1k", func(b *testing.B) {

		stack := &LinkedListStack{}

		b.ResetTimer()

		for range b.N {
			stack.Pop()
		}

	})

	b.Run("PopFrom10k", func(b *testing.B) {

		stack := &LinkedListStack{}

		b.ResetTimer()

		for range b.N {
			stack.Pop()
		}

	})

}
