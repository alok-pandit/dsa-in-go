package arraystack

import (
	"fmt"
	"testing"
)

type TestCase struct {
	name     string
	capacity int
	pushData []interface{}
	wantTop  int
	wantLen  int
}

func TestNewArrayStack(t *testing.T) {
	tests := []TestCase{
		{
			name:     "Zero capacity",
			capacity: 0,
			pushData: []interface{}{1},
			wantTop:  -1,
			wantLen:  0,
		},
		{
			name:     "Positive capacity",
			capacity: 5,
			pushData: []interface{}{1},
			wantTop:  -1,
			wantLen:  0,
		},
		{
			name:     "Large capacity",
			capacity: 1000000,
			pushData: []interface{}{1},
			wantTop:  -1,
			wantLen:  0,
		}}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {

				stack := NewArrayStack(tt.capacity)

				if stack == nil {
					t.Fatal("NewArrayStack returned nil")
				}

				if cap(stack.data) != tt.capacity {
					t.Errorf("NewArrayStack(%d) data capacity = %d, want %d", tt.capacity, cap(stack.data), tt.capacity)
				}

				if stack.top != tt.wantTop {
					t.Errorf("NewArrayStack(%d) top = %d, want %d", tt.capacity, stack.top, tt.wantTop)
				}

				if stack.length != tt.wantLen {
					t.Errorf("NewArrayStack(%d) length = %d, want %d", tt.capacity, stack.length, tt.wantLen)
				}

			},
		)
	}
}

func TestNewArrayStackNegativeCapacity(t *testing.T) {
	stack := NewArrayStack(-1)
	if stack == nil {
		t.Fatal("NewArrayStack returned nil for negative capacity")
	}
	if cap(stack.data) != 0 {
		t.Errorf("NewArrayStack(-1) data capacity = %d, want 1024", len(stack.data))
	}
	if stack.top != -1 {
		t.Errorf("NewArrayStack(-1) top = %d, want -1", stack.top)
	}
	if stack.length != 0 {
		t.Errorf("NewArrayStack(-1) length = %d, want 0", stack.length)
	}
}
func TestArrayStackPush(t *testing.T) {
	tests := []TestCase{
		{
			name:     "Push to empty stack",
			capacity: 5,
			pushData: []interface{}{1},
			wantTop:  0,
			wantLen:  1,
		},
		{
			name:     "Push multiple items",
			capacity: 5,
			pushData: []interface{}{1, 2, 3},
			wantTop:  2,
			wantLen:  3,
		},
		{
			name:     "Push to full stack",
			capacity: 3,
			pushData: []interface{}{1, 2, 3, 4},
			wantTop:  3,
			wantLen:  4,
		},
		{
			name:     "Push different data types",
			capacity: 5,
			pushData: []interface{}{1, "two", 3.14, true},
			wantTop:  3,
			wantLen:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewArrayStack(tt.capacity)
			for _, data := range tt.pushData {
				stack.Push(data)
			}
			fmt.Printf("%+v: %+v CAP: %+v", tt.name, stack, cap(stack.data))

			if stack.top != tt.wantTop {
				t.Errorf("Push() top = %d, want %d", stack.top, tt.wantTop)
			}
			if stack.length != tt.wantLen {
				t.Errorf("Push() length = %d, want %d", stack.length, tt.wantLen)
			}
			for i := 0; i <= tt.wantTop; i++ {
				if stack.data[i] != tt.pushData[i] {
					t.Errorf("Push() data[%d] = %v, want %v", i, stack.data[i], tt.pushData[i])
				}
			}
		})
	}
}
func TestArrayStackPop(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		pushData []interface{}
		wantPop  []interface{}
		wantTop  int
		wantLen  int
	}{
		{
			name:     "Pop from empty stack",
			capacity: 5,
			pushData: []interface{}{},
			wantPop:  []interface{}{nil},
			wantTop:  -1,
			wantLen:  0,
		},
		{
			name:     "Pop single item",
			capacity: 5,
			pushData: []interface{}{1},
			wantPop:  []interface{}{1, nil},
			wantTop:  -1,
			wantLen:  0,
		},
		{
			name:     "Pop multiple items",
			capacity: 5,
			pushData: []interface{}{1, 2, 3},
			wantPop:  []interface{}{3, 2, 1, nil},
			wantTop:  -1,
			wantLen:  0,
		},
		{
			name:     "Pop different data types",
			capacity: 5,
			pushData: []interface{}{1, "two", 3.14, true},
			wantPop:  []interface{}{true, 3.14, "two", 1, nil},
			wantTop:  -1,
			wantLen:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewArrayStack(tt.capacity)
			for _, data := range tt.pushData {
				stack.Push(data)
			}

			fmt.Printf("%+v: %+v CAP: %+v TOP: %+v\n\n", tt.name, stack, cap(stack.data), stack.top)

			for _, want := range tt.wantPop {
				got := stack.Pop()

				fmt.Printf("Pop: Got: %v Want: %+v\n", want, got)

				if got != want {
					t.Errorf("Pop() want: %v got: %v", want, got)
				}
			}

			fmt.Println(stack.top, tt.wantTop, stack.length, tt.wantLen)
			if stack.top != tt.wantTop {
				t.Errorf("After Pop() top = %d, want %d", stack.top, tt.wantTop)
			}
			if stack.length != tt.wantLen {
				t.Errorf("After Pop() length = %d, want %d", stack.length, tt.wantLen)
			}
		})
	}
}

func TestArrayStackPopPushCombination(t *testing.T) {
	stack := NewArrayStack(5)

	// Push and pop alternately
	stack.Push(1)
	if got := stack.Pop(); got != 1 {
		t.Errorf("Pop() after Push(1) = %v, want 1", got)
	}

	stack.Push(2)
	stack.Push(3)
	if got := stack.Pop(); got != 3 {
		t.Errorf("Pop() after Push(2) and Push(3) = %v, want 3", got)
	}

	stack.Push(4)
	if got := stack.Pop(); got != 4 {
		t.Errorf("Pop() after Push(4) = %v, want 4", got)
	}
	if got := stack.Pop(); got != 2 {
		t.Errorf("Pop() after previous Pop() = %v, want 2", got)
	}

	// Push to fill the stack and then pop all
	for i := 1; i <= 5; i++ {
		stack.Push(i)
	}

	for i := 5; i >= 1; i-- {
		if got := stack.Pop(); got != i {
			t.Errorf("Pop() %d = %v, want %d", i, got, i)
		}
	}

	// Ensure the stack is empty
	if got := stack.Pop(); got != nil {
		t.Errorf("Pop() from empty stack = %v, want nil", got)
	}

	if stack.top != -1 || stack.length != 0 {
		t.Errorf("After all operations, top = %d, length = %d, want top = -1, length = 0", stack.top, stack.length)
	}
}

func TestArrayStackPeek(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		pushData []interface{}
		wantPeek interface{}
	}{
		{
			name:     "Peek empty stack",
			capacity: 5,
			pushData: []interface{}{},
			wantPeek: nil,
		},
		{
			name:     "Peek single item",
			capacity: 5,
			pushData: []interface{}{1},
			wantPeek: 1,
		},
		{
			name:     "Peek multiple items",
			capacity: 5,
			pushData: []interface{}{1, 2, 3},
			wantPeek: 3,
		},
		{
			name:     "Peek different data types",
			capacity: 5,
			pushData: []interface{}{1, "two", 3.14, true},
			wantPeek: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewArrayStack(tt.capacity)
			for _, data := range tt.pushData {
				stack.Push(data)
			}

			got := stack.Peek()
			if got != tt.wantPeek {
				t.Errorf("Peek() = %v, want %v", got, tt.wantPeek)
			}

			if len(tt.pushData) > 0 {
				if stack.top != len(tt.pushData)-1 {
					t.Errorf("After Peek() top = %d, want %d", stack.top, len(tt.pushData)-1)
				}
				if stack.length != len(tt.pushData) {
					t.Errorf("After Peek() length = %d, want %d", stack.length, len(tt.pushData))
				}
			} else {
				if stack.top != -1 {
					t.Errorf("After Peek() on empty stack, top = %d, want -1", stack.top)
				}
				if stack.length != 0 {
					t.Errorf("After Peek() on empty stack, length = %d, want 0", stack.length)
				}
			}
		})
	}
}

func TestArrayStackPeekAfterPushPop(t *testing.T) {
	stack := NewArrayStack(5)

	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	got := stack.Peek()
	if got != 1 {
		t.Errorf("Peek() after Push(1), Push(2), Pop() = %v, want 1", got)
	}

	stack.Push(3)
	got = stack.Peek()
	if got != 3 {
		t.Errorf("Peek() after Push(3) = %v, want 3", got)
	}

	stack.Pop()
	stack.Pop()
	got = stack.Peek()
	if got != nil {
		t.Errorf("Peek() after popping all items = %v, want nil", got)
	}
}

func TestArrayStackPeekDoesNotModifyStack(t *testing.T) {
	stack := NewArrayStack(5)
	stack.Push(1)
	stack.Push(2)

	initialTop := stack.top
	initialLength := stack.length

	for i := 0; i < 3; i++ {
		got := stack.Peek()
		if got != 2 {
			t.Errorf("Peek() iteration %d = %v, want 2", i, got)
		}
		if stack.top != initialTop {
			t.Errorf("After Peek() iteration %d, top changed: got %d, want %d", i, stack.top, initialTop)
		}
		if stack.length != initialLength {
			t.Errorf("After Peek() iteration %d, length changed: got %d, want %d", i, stack.length, initialLength)
		}
	}
}
