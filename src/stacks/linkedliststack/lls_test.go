package linkedliststack

import (
	"testing"
)

func TestPush(t *testing.T) {
	tests := []struct {
		name     string
		data     interface{}
		expected int
	}{
		{"Push integer", 42, 1},
		{"Push string", "hello", 1},
		{"Push struct", struct{ value int }{10}, 1},
		{"Push nil", nil, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LinkedListStack{}
			s.Push(tt.data)

			if s.Size() != tt.expected {
				t.Errorf("Push() size = %v, want %v", s.Size(), tt.expected)
			}

			if s.top == nil {
				t.Errorf("Push() top is nil, want non-nil")
			}

			if s.top.data != tt.data {
				t.Errorf("Push() top data = %v, want %v", s.top.data, tt.data)
			}

			if s.top.prev != nil {
				t.Errorf("Push() top.prev is non-nil, want nil")
			}
		})
	}
}

func TestPushMultiple(t *testing.T) {

	s := &LinkedListStack{}

	data := []interface{}{1, "two", 3.14, true}

	for i, item := range data {
		s.Push(item)
		if s.Size() != i+1 {
			t.Errorf("Push() size = %v, want %v", s.Size(), i+1)
		}
	}

	if s.Size() != len(data) {
		t.Errorf("Push() final size = %v, want %v", s.Size(), len(data))
	}

	for i := len(data) - 1; i >= 0; i-- {
		if s.top.data != data[i] {
			t.Errorf("Push() top data = %v, want %v", s.top.data, data[i])
		}
		s.top = s.top.prev
	}

	if s.top != nil {
		t.Errorf("Push() after traversal, top is non-nil, want nil")
	}

}

func TestPushLargeNumber(t *testing.T) {
	s := &LinkedListStack{}

	n := 1000000

	for i := 0; i < n; i++ {
		s.Push(i)
	}

	if s.Size() != n {
		t.Errorf("Push() size = %v, want %v", s.Size(), n)
	}

	for i := n - 1; i >= 0; i-- {
		if s.top.data != i {
			t.Errorf("Push() top data = %v, want %v", s.top.data, i)
		}
		s.top = s.top.prev
	}

	if s.top != nil {
		t.Errorf("Push() after traversal, top is non-nil, want nil")
	}

}

func TestPop(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *LinkedListStack
		expected interface{}
		wantSize int
	}{
		{
			name: "Pop from empty stack",
			setup: func() *LinkedListStack {
				return &LinkedListStack{}
			},
			expected: nil,
			wantSize: 0,
		},
		{
			name: "Pop single element",
			setup: func() *LinkedListStack {
				s := &LinkedListStack{}
				s.Push(42)
				return s
			},
			expected: 42,
			wantSize: 0,
		},
		{
			name: "Pop multiple elements",
			setup: func() *LinkedListStack {
				s := &LinkedListStack{}
				s.Push(1)
				s.Push(2)
				s.Push(3)
				return s
			},
			expected: 3,
			wantSize: 2,
		},
		{
			name: "Pop with different data types",
			setup: func() *LinkedListStack {
				s := &LinkedListStack{}
				s.Push("hello")
				s.Push(42)
				s.Push(true)
				return s
			},
			expected: true,
			wantSize: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.setup()
			result := s.Pop()

			if result != tt.expected {
				t.Errorf("Pop() = %v, want %v", result, tt.expected)
			}

			if s.Size() != tt.wantSize {
				t.Errorf("Size after Pop() = %v, want %v", s.Size(), tt.wantSize)
			}

			if tt.wantSize == 0 && s.top != nil {
				t.Errorf("Top after Pop() is non-nil, want nil")
			}
		})
	}
}

func TestPopUntilEmpty(t *testing.T) {

	s := &LinkedListStack{}

	elements := []interface{}{1, "two", 3.14, true}

	for _, elem := range elements {
		s.Push(elem)
	}

	for i := len(elements) - 1; i >= 0; i-- {
		popped := s.Pop()
		if popped != elements[i] {
			t.Errorf("Pop() = %v, want %v", popped, elements[i])
		}
	}

	if s.Size() != 0 {
		t.Errorf("Size after popping all elements = %v, want 0", s.Size())
	}

	if s.top != nil {
		t.Errorf("Top after popping all elements is non-nil, want nil")
	}

	lastPop := s.Pop()

	if lastPop != nil {
		t.Errorf("Pop() from empty stack = %v, want nil", lastPop)
	}

}

func TestPopPerformance(t *testing.T) {

	s := &LinkedListStack{}

	n := 1000000

	for i := 0; i < n; i++ {
		s.Push(i)
	}

	for i := n - 1; i >= 0; i-- {
		popped := s.Pop()
		if popped != i {
			t.Errorf("Pop() = %v, want %v", popped, i)
		}
	}

	if s.Size() != 0 {
		t.Errorf("Size after popping all elements = %v, want 0", s.Size())
	}

	if s.top != nil {
		t.Errorf("Top after popping all elements is non-nil, want nil")
	}

}
func TestPeek(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *LinkedListStack
		expected interface{}
	}{
		{
			name: "Peek empty stack",
			setup: func() *LinkedListStack {
				return &LinkedListStack{}
			},
			expected: nil,
		},
		{
			name: "Peek single element",
			setup: func() *LinkedListStack {
				s := &LinkedListStack{}
				s.Push(42)
				return s
			},
			expected: 42,
		},
		{
			name: "Peek multiple elements",
			setup: func() *LinkedListStack {
				s := &LinkedListStack{}
				s.Push(1)
				s.Push(2)
				s.Push(3)
				return s
			},
			expected: 3,
		},
		{
			name: "Peek after push and pop",
			setup: func() *LinkedListStack {
				s := &LinkedListStack{}
				s.Push(1)
				s.Push(2)
				s.Pop()
				return s
			},
			expected: 1,
		},
		{
			name: "Peek with different data types",
			setup: func() *LinkedListStack {
				s := &LinkedListStack{}
				s.Push("hello")
				s.Push(42)
				s.Push(true)
				return s
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.setup()
			result := s.Peek()

			if result != tt.expected {
				t.Errorf("Peek() = %v, want %v", result, tt.expected)
			}

		})

	}

}

func TestPeekConsistency(t *testing.T) {
	s := &LinkedListStack{}
	elements := []interface{}{1, "two", 3.14, true, nil}

	for _, elem := range elements {
		s.Push(elem)
		peeked := s.Peek()
		if peeked != elem {
			t.Errorf("Peek() after Push(%v) = %v, want %v", elem, peeked, elem)
		}
	}

	for i := len(elements) - 1; i >= 0; i-- {
		peeked := s.Peek()
		if peeked != elements[i] {
			t.Errorf("Peek() before Pop() = %v, want %v", peeked, elements[i])
		}
		popped := s.Pop()
		if popped != elements[i] {
			t.Errorf("Pop() = %v, want %v", popped, elements[i])
		}
	}

	if s.Peek() != nil {
		t.Errorf("Peek() on empty stack = %v, want nil", s.Peek())
	}
}

func TestPeekPerformance(t *testing.T) {
	s := &LinkedListStack{}
	n := 1000000

	for i := 0; i < n; i++ {
		s.Push(i)
	}

	for i := 0; i < n; i++ {
		peeked := s.Peek()
		if peeked != n-1 {
			t.Errorf("Peek() = %v, want %v", peeked, n-1)
		}
	}

	if s.Size() != n {
		t.Errorf("Size after multiple Peek() calls = %v, want %v", s.Size(), n)
	}
}
