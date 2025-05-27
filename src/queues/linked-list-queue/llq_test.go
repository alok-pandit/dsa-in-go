package linkedlistqueue

import (
	"testing"
)

type TestCase struct {
	name     string
	data     []string
	expected []string
}

func TestLinkedListQueue(t *testing.T) {
	tests := []TestCase{
		{
			name:     "Basic queue operations",
			data:     []string{"first", "second", "third"},
			expected: []string{"first", "second", "third"},
		},
		{
			name:     "Empty queue operations",
			data:     []string{},
			expected: []string{},
		},
		{
			name:     "Single element operations",
			data:     []string{"only"},
			expected: []string{"only"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			queue := &LinkedListQueue{}

			// Test IsEmpty on new queue
			if !queue.IsEmpty() {
				t.Error("New queue should be empty")
			}

			// Test Enqueue
			for _, item := range test.data {
				queue.Enqueue(item)
			}

			// Test Front
			if len(test.data) > 0 {
				front, err := queue.Front()
				if err != nil {
					t.Errorf("Front() error = %v", err)
				}
				if front != test.data[0] {
					t.Errorf("Front() = %v, want %v", front, test.data[0])
				}
			} else {
				_, err := queue.Front()
				if err == nil {
					t.Error("Front() should return error for empty queue")
				}
			}

			// Test Dequeue
			for i, want := range test.expected {
				got, err := queue.Dequeue()
				if err != nil {
					t.Errorf("Dequeue() error = %v", err)
				}
				if got != want {
					t.Errorf("Dequeue() [%d] = %v, want %v", i, got, want)
				}
			}

			// Test Dequeue on empty queue
			_, err := queue.Dequeue()
			if err == nil {
				t.Error("Dequeue() should return error for empty queue")
			}
		})
	}
}

func TestQueueOperations(t *testing.T) {
	queue := &LinkedListQueue{}

	// Test empty queue operations
	if !queue.IsEmpty() {
		t.Error("New queue should be empty")
	}

	_, err := queue.Front()
	if err == nil {
		t.Error("Front() should return error for empty queue")
	}

	_, err = queue.Dequeue()
	if err == nil {
		t.Error("Dequeue() should return error for empty queue")
	}

	// Test mixed operations
	queue.Enqueue("first")
	queue.Enqueue("second")

	if queue.IsEmpty() {
		t.Error("Queue should not be empty after enqueue")
	}

	front, err := queue.Front()
	if err != nil || front != "first" {
		t.Errorf("Front() = %v, want 'first'", front)
	}

	queue.Enqueue("third")
	dequeued, _ := queue.Dequeue()
	if dequeued != "first" {
		t.Errorf("Dequeue() = %v, want 'first'", dequeued)
	}

	front, _ = queue.Front()
	if front != "second" {
		t.Errorf("Front() after dequeue = %v, want 'second'", front)
	}
}
func BenchmarkEnqueue(b *testing.B) {
	queue := &LinkedListQueue{}
	b.ResetTimer()
	for i := range b.N {
		queue.Enqueue(i)
	}
}
func BenchmarkDequeue(b *testing.B) {
	queue := &LinkedListQueue{}
	i := 0
	for i = range b.N {
		queue.Enqueue(i)
	}
	b.ResetTimer()
	for range b.N {
		queue.Dequeue()
	}
}

func BenchmarkFront(b *testing.B) {
	queue := &LinkedListQueue{}
	queue.Enqueue("test")
	b.ResetTimer()
	for range b.N {
		queue.Front()
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	queue := &LinkedListQueue{}
	b.ResetTimer()
	for range b.N {
		queue.IsEmpty()
	}
}
