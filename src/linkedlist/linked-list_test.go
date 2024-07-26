package linkedlist

import (
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
