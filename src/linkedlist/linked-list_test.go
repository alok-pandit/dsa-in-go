package linkedlist

import (
	"strconv"
	"testing"
)

type TestCase struct {
	name     string
	data     []string
	nodeData string
	newData  string
	want     []string
}

func testValidator(l *SinglyLinkedList, t *testing.T, test TestCase) {

	l.PrintList()

	var got []string

	currentNode := l.head

	for currentNode != nil {
		got = append(got, currentNode.data)
		currentNode = currentNode.next
	}

	if len(got) != len(test.want) {
		t.Errorf("got %v, want %v", got, test.want)
	}

	for i := range got {
		if got[i] != test.want[i] {
			t.Errorf("got %v, want %v", got, test.want)
			break
		}
	}

}

func getInitialList(test TestCase) *SinglyLinkedList {

	l := &SinglyLinkedList{}

	for _, d := range test.data {
		l.AppendToEnd(d)
	}

	return l

}

func TestAppendBefore(t *testing.T) {

	tests := []TestCase{
		{
			name:     "Insert at the beginning",
			data:     []string{"P", "Q", "R"},
			nodeData: "P",
			newData:  "X",
			want:     []string{"X", "P", "Q", "R"},
		},
		{
			name:     "Insert in the middle",
			data:     []string{"A", "B", "C"},
			nodeData: "B",
			newData:  "X",
			want:     []string{"A", "X", "B", "C"},
		},
		{
			name:     "Insert with an empty list",
			data:     []string{},
			nodeData: "",
			newData:  "X",
			want:     []string{"X"},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			l := getInitialList(test)

			l.AppendBefore(test.newData, test.nodeData)

			testValidator(l, t, test)

		})
	}
}

func TestAppendAfter(t *testing.T) {

	tests := []TestCase{
		{
			name:     "Insert at the End",
			data:     []string{"P", "Q", "R"},
			nodeData: "R",
			newData:  "X",
			want:     []string{"P", "Q", "R", "X"},
		},
		{
			name:     "Insert in the middle",
			data:     []string{"A", "B", "C"},
			nodeData: "B",
			newData:  "X",
			want:     []string{"A", "B", "X", "C"},
		},
		{
			name:     "Insert in the middle in presence of dupes",
			data:     []string{"A", "B", "B", "B", "C"},
			nodeData: "B",
			newData:  "X",
			want:     []string{"A", "B", "X", "B", "B", "C"},
		},
		{
			name:     "Insert with an empty list",
			data:     []string{},
			nodeData: "",
			newData:  "X",
			want:     []string{"X"},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			l := getInitialList(test)

			l.AppendAfter(test.newData, test.nodeData)

			testValidator(l, t, test)

		})
	}
}

func TestReverse(t *testing.T) {

	tests := []TestCase{
		{
			name: "Reverse empty list",
			data: []string{},
			want: []string{},
		},
		{
			name: "Reverse single node",
			data: []string{"A"},
			want: []string{"A"},
		},
		{
			name: "Reverse multiple nodes",
			data: []string{"A", "B", "C", "D"},
			want: []string{"D", "C", "B", "A"},
		},
		{
			name: "Reverse multiple nodes with duplicate values",
			data: []string{"A", "B", "C", "B", "A"},
			want: []string{"A", "B", "C", "B", "A"},
		},
		{
			name: "Reverse multiple nodes with consecutive duplicate values",
			data: []string{"A", "B", "B", "C", "C", "B", "A"},
			want: []string{"A", "B", "C", "C", "B", "B", "A"},
		},
		{
			name: "Reverse multiple nodes with negative values",
			data: []string{"1", "-2", "3", "-4"},
			want: []string{"-4", "3", "-2", "1"},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			l := getInitialList(test)

			l.Reverse()

			testValidator(l, t, test)

		})
	}
}

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

func BenchmarkReverse10k(b *testing.B) {
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

func BenchmarkAddToStart(b *testing.B) {
	// Create a sample linked list with a specific size
	list := &SinglyLinkedList{}
	for i := range listLength {
		list.AppendBeforeStart(strconv.Itoa(i))
	}

	// Benchmark the Reverse function
	b.ResetTimer()
	list.AppendBeforeStart("A")

}

func BenchmarkAddToEnd(b *testing.B) {
	// Create a sample linked list with a specific size
	list := &SinglyLinkedList{}
	for i := range listLength {
		list.AppendBeforeStart(strconv.Itoa(i))
	}

	// Benchmark the Reverse function
	b.ResetTimer()
	list.AppendToEnd("A")

}
