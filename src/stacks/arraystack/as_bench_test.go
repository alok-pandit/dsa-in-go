package arraystack

import (
	"testing"
)

func BenchmarkPush(b *testing.B) {

	b.Run("PushToEmptyStack", func(b *testing.B) {

		stack := NewArrayStack(5)

		b.ResetTimer()

		for i := range b.N {
			stack.Push(i)
		}

	})

	b.Run("PushAtop2kFilled", func(b *testing.B) {

		stack := NewArrayStack(5)

		for i := range 2000 {
			stack.Push(i)
		}

		b.ResetTimer()

		for i := range b.N {
			stack.Push(i)
		}

	})

	b.Run("PushTo2kCapEmpty", func(b *testing.B) {

		stack := NewArrayStack(2000)

		b.ResetTimer()

		for i := range b.N {
			stack.Push(i)
		}

	})

}

func BenchmarkPop(b *testing.B) {

	b.Run("PopFromEmptyStack", func(b *testing.B) {

		stack := NewArrayStack(5)

		b.ResetTimer()

		for range b.N {
			stack.Pop()
		}

	})

	b.Run("PopFrom1k", func(b *testing.B) {

		stack := NewArrayStack(1000)

		b.ResetTimer()

		for range b.N {
			stack.Pop()
		}

	})

	b.Run("PopFrom10k", func(b *testing.B) {

		stack := NewArrayStack(10000)

		b.ResetTimer()

		for range b.N {
			stack.Pop()
		}

	})

}
