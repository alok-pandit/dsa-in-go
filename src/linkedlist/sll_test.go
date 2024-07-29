package linkedlist

import (
	"bytes"
	"io"
	"os"
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

func TestDeleteHead(t *testing.T) {
	tests := []TestCase{
		{
			name: "Delete head from list with multiple elements",
			data: []string{"A", "B", "C"},
			want: []string{"B", "C"},
		},
		{
			name: "Delete head from list with one element",
			data: []string{"A"},
			want: []string{},
		},
		{
			name: "Delete head from empty list",
			data: []string{},
			want: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			l.DeleteHead()
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestDeleteTail(t *testing.T) {
	tests := []TestCase{
		{
			name: "Delete tail from list with multiple elements",
			data: []string{"A", "B", "C"},
			want: []string{"A", "B"},
		},
		{
			name: "Delete tail from list with two elements",
			data: []string{"A", "B"},
			want: []string{"A"},
		},
		{
			name: "Delete tail from list with one element",
			data: []string{"A"},
			want: []string{},
		},
		{
			name: "Delete tail from empty list",
			data: []string{},
			want: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			l.DeleteTail()
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestDeleteNode(t *testing.T) {
	tests := []TestCase{
		{
			name:     "Delete node from middle of list",
			data:     []string{"A", "B", "C", "D"},
			nodeData: "C",
			want:     []string{"A", "B", "D"},
		},
		{
			name:     "Delete first node",
			data:     []string{"A", "B", "C"},
			nodeData: "A",
			want:     []string{"B", "C"},
		},
		{
			name:     "Delete last node",
			data:     []string{"A", "B", "C"},
			nodeData: "C",
			want:     []string{"A", "B"},
		},
		{
			name:     "Delete node from list with one element",
			data:     []string{"A"},
			nodeData: "A",
			want:     []string{},
		},
		{
			name:     "Delete non-existent node",
			data:     []string{"A", "B", "C"},
			nodeData: "D",
			want:     []string{"A", "B", "C"},
		},
		{
			name:     "Delete from empty list",
			data:     []string{},
			nodeData: "A",
			want:     []string{},
		},
		{
			name:     "Delete node with duplicate values",
			data:     []string{"A", "B", "C", "B", "D"},
			nodeData: "B",
			want:     []string{"A", "C", "B", "D"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(test)
			l.DeleteNode(test.nodeData)
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestDeleteOneBefore(t *testing.T) {
	tests := []TestCase{
		{
			name:     "Delete node before middle element",
			data:     []string{"A", "B", "C", "D", "E"},
			nodeData: "D",
			want:     []string{"A", "B", "D", "E"},
		},
		{
			name:     "Delete node before last element",
			data:     []string{"A", "B", "C", "D"},
			nodeData: "D",
			want:     []string{"A", "B", "D"},
		},
		{
			name:     "Delete node before first element (no change)",
			data:     []string{"A", "B", "C"},
			nodeData: "A",
			want:     []string{"A", "B", "C"},
		},
		{
			name:     "Delete from list with two elements",
			data:     []string{"A", "B"},
			nodeData: "B",
			want:     []string{"B"},
		},
		{
			name:     "Delete from list with one element (no change)",
			data:     []string{"A"},
			nodeData: "A",
			want:     []string{"A"},
		},
		{
			name:     "Delete from empty list (no change)",
			data:     []string{},
			nodeData: "A",
			want:     []string{},
		},
		{
			name:     "Delete before non-existent node (no change)",
			data:     []string{"A", "B", "C"},
			nodeData: "D",
			want:     []string{"A", "B", "C"},
		},
		{
			name:     "Delete before duplicate node (first occurrence)",
			data:     []string{"A", "B", "C", "B", "D"},
			nodeData: "B",
			want:     []string{"B", "C", "B", "D"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(test)
			l.DeleteOneBefore(test.nodeData)
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestDeleteOneAfter(t *testing.T) {
	tests := []TestCase{
		{
			name:     "Delete node after middle element",
			data:     []string{"A", "B", "C", "D", "E"},
			nodeData: "B",
			want:     []string{"A", "B", "D", "E"},
		},
		{
			name:     "Delete last node",
			data:     []string{"A", "B", "C", "D"},
			nodeData: "C",
			want:     []string{"A", "B", "C"},
		},
		{
			name:     "Delete after first element",
			data:     []string{"A", "B", "C"},
			nodeData: "A",
			want:     []string{"A", "C"},
		},
		{
			name:     "Delete from list with two elements",
			data:     []string{"A", "B"},
			nodeData: "A",
			want:     []string{"A"},
		},
		{
			name:     "Delete from list with one element (no change)",
			data:     []string{"A"},
			nodeData: "A",
			want:     []string{"A"},
		},
		{
			name:     "Delete from empty list (no change)",
			data:     []string{},
			nodeData: "A",
			want:     []string{},
		},
		{
			name:     "Delete after non-existent node (no change)",
			data:     []string{"A", "B", "C"},
			nodeData: "D",
			want:     []string{"A", "B", "C"},
		},
		{
			name:     "Delete after duplicate node (first occurrence)",
			data:     []string{"A", "B", "C", "B", "D", "E"},
			nodeData: "B",
			want:     []string{"A", "B", "B", "D", "E"},
		},
		{
			name:     "Delete after consecutive duplicate node (first occurrence)",
			data:     []string{"A", "B", "B", "D", "E"},
			nodeData: "B",
			want:     []string{"A", "B", "D", "E"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(test)
			l.DeleteOneAfter(test.nodeData)
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestDeleteAllBefore(t *testing.T) {
	tests := []TestCase{
		{
			name:     "Delete all before middle element",
			data:     []string{"A", "B", "C", "D", "E"},
			nodeData: "C",
			want:     []string{"C", "D", "E"},
		},
		{
			name:     "Delete all before last element",
			data:     []string{"A", "B", "C", "D"},
			nodeData: "D",
			want:     []string{"D"},
		},
		{
			name:     "Delete all before first element (no change)",
			data:     []string{"A", "B", "C"},
			nodeData: "A",
			want:     []string{"A", "B", "C"},
		},
		{
			name:     "Delete all before non-existent node (no change)",
			data:     []string{"A", "B", "C"},
			nodeData: "D",
			want:     []string{"A", "B", "C"},
		},
		{
			name:     "Delete all before in list with duplicate values",
			data:     []string{"A", "B", "C", "B", "D", "E"},
			nodeData: "B",
			want:     []string{"B", "C", "B", "D", "E"},
		},
		{
			name:     "Delete all before in empty list",
			data:     []string{},
			nodeData: "A",
			want:     []string{},
		},
		{
			name:     "Delete all before in list with one element",
			data:     []string{"A"},
			nodeData: "A",
			want:     []string{"A"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(test)
			l.DeleteAllBefore(test.nodeData)
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestDeleteAllAfter(t *testing.T) {
	tests := []TestCase{
		{
			name:     "Delete all after middle element",
			data:     []string{"A", "B", "C", "D", "E"},
			nodeData: "C",
			want:     []string{"A", "B", "C"},
		},
		{
			name:     "Delete all after first element",
			data:     []string{"A", "B", "C", "D"},
			nodeData: "A",
			want:     []string{"A"},
		},
		{
			name:     "Delete all after last element (no change)",
			data:     []string{"A", "B", "C"},
			nodeData: "C",
			want:     []string{"A", "B", "C"},
		},
		{
			name:     "Delete all after non-existent node (no change)",
			data:     []string{"A", "B", "C"},
			nodeData: "D",
			want:     []string{"A", "B", "C"},
		},
		{
			name:     "Delete all after in list with duplicate values",
			data:     []string{"A", "B", "C", "B", "D", "E"},
			nodeData: "B",
			want:     []string{"A", "B"},
		},
		{
			name:     "Delete all after in empty list",
			data:     []string{},
			nodeData: "A",
			want:     []string{},
		},
		{
			name:     "Delete all after in list with one element",
			data:     []string{"A"},
			nodeData: "A",
			want:     []string{"A"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(test)
			l.DeleteAllAfter(test.nodeData)
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestPrintList(t *testing.T) {
	tests := []struct {
		name string
		data []string
		want string
	}{
		{
			name: "Print empty list",
			data: []string{},
			want: "\nList is empty\n",
		},
		{
			name: "Print list with one element",
			data: []string{"A"},
			want: "\nA\n",
		},
		{
			name: "Print list with multiple elements",
			data: []string{"A", "B", "C"},
			want: "\nA -> B -> C\n",
		},
		{
			name: "Print list with duplicate elements",
			data: []string{"A", "B", "B", "C"},
			want: "\nA -> B -> B -> C\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: tt.data})

			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			l.PrintList()

			w.Close()
			os.Stdout = oldStdout

			var buf bytes.Buffer
			io.Copy(&buf, r)
			got := buf.String()

			if got != tt.want {
				t.Errorf("PrintList() = %q, want %q", got, tt.want)
			}
		})
	}
}
