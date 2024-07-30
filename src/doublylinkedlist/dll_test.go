package doublylinkedlist

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

func testValidator(l *DoublyLinkedList, t *testing.T, test TestCase) {

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

func getInitialList(test TestCase) *DoublyLinkedList {

	l := &DoublyLinkedList{}

	for _, d := range test.data {
		l.Append(d)
	}

	return l

}

func TestAppend(t *testing.T) {
	tests := []TestCase{
		{
			name:     "Append to empty list",
			data:     []string{},
			newData:  "first",
			nodeData: "",
			want:     []string{"first"},
		},
		{
			name:     "Append to list with one element",
			data:     []string{"first"},
			newData:  "second",
			nodeData: "",
			want:     []string{"first", "second"},
		},
		{
			name:     "Append to list with multiple elements",
			data:     []string{"first", "second", "third"},
			newData:  "fourth",
			nodeData: "",
			want:     []string{"first", "second", "third", "fourth"},
		},
		{
			name:     "Append empty string",
			data:     []string{"first", "second"},
			newData:  "",
			nodeData: "",
			want:     []string{"first", "second", ""},
		},
	}

	for _, test := range tests {

		t.Run(
			test.name,
			func(t *testing.T) {

				l := getInitialList(test)

				l.Append(test.newData)

				testValidator(l, t, test)

			})

	}

}
func TestPrepend(t *testing.T) {
	tests := []TestCase{
		{
			name:    "Prepend to empty list",
			data:    []string{},
			newData: "first",
			want:    []string{"first"},
		},
		{
			name:    "Prepend to list with one element",
			data:    []string{"second"},
			newData: "first",
			want:    []string{"first", "second"},
		},
		{
			name:    "Prepend to list with multiple elements",
			data:    []string{"second", "third", "fourth"},
			newData: "first",
			want:    []string{"first", "second", "third", "fourth"},
		},
		{
			name:    "Prepend empty string",
			data:    []string{"first", "second"},
			newData: "",
			want:    []string{"", "first", "second"},
		},
	}

	for _, test := range tests {
		t.Run(
			test.name,
			func(t *testing.T) {

				l := getInitialList(test)

				l.Prepend(test.newData)

				testValidator(l, t, test)

			})

	}

}
func TestFindNode(t *testing.T) {
	tests := []TestCase{
		{
			name:     "Find node in list with single element",
			data:     []string{"first"},
			nodeData: "first",
			want:     []string{"first"},
		},
		{
			name:     "Find node in list with multiple elements",
			data:     []string{"first", "second", "third"},
			nodeData: "second",
			want:     []string{"second"},
		},
		{
			name:     "Find last node in list",
			data:     []string{"first", "second", "third"},
			nodeData: "third",
			want:     []string{"third"},
		},
		{
			name:     "Node not found in list",
			data:     []string{"first", "second", "third"},
			nodeData: "fourth",
			want:     nil,
		},
		{
			name:     "Find node in empty list",
			data:     []string{},
			nodeData: "first",
			want:     nil,
		},
		{
			name:     "Find empty string node",
			data:     []string{"first", "", "third"},
			nodeData: "",
			want:     []string{""},
		},
	}

	for _, test := range tests {
		t.Run(
			test.name,
			func(t *testing.T) {

				l := getInitialList(test)

				got := l.findNode(test.nodeData)

				list := &DoublyLinkedList{}

				if got != nil {
					list.Append(got.data)
				}

				testValidator(list, t, test)

			},
		)
	}
}

func TestAppendAfter(t *testing.T) {
	tests := []TestCase{
		{
			name:     "Append after first element",
			data:     []string{"first", "third"},
			nodeData: "first",
			newData:  "second",
			want:     []string{"first", "second", "third"},
		},
		{
			name:     "Append after last element",
			data:     []string{"first", "second"},
			nodeData: "second",
			newData:  "third",
			want:     []string{"first", "second", "third"},
		},
		{
			name:     "Append after middle element",
			data:     []string{"first", "second", "fourth"},
			nodeData: "second",
			newData:  "third",
			want:     []string{"first", "second", "third", "fourth"},
		},
		{
			name:     "Append after non-existent element",
			data:     []string{"first", "second"},
			nodeData: "third",
			newData:  "fourth",
			want:     []string{"first", "second"},
		},
		{
			name:     "Append empty string after element",
			data:     []string{"first", "second"},
			nodeData: "first",
			newData:  "",
			want:     []string{"first", "", "second"},
		},
		{
			name:     "Append after empty string element",
			data:     []string{"first", "", "third"},
			nodeData: "",
			newData:  "second",
			want:     []string{"first", "", "second", "third"},
		},
	}

	for _, test := range tests {

		t.Run(
			test.name,
			func(t *testing.T) {

				l := getInitialList(test)

				l.AppendAfter(test.nodeData, test.newData)

				testValidator(l, t, test)

			},
		)

	}

}
func TestAppendBefore(t *testing.T) {
	tests := []TestCase{
		{
			name:     "Append before first element",
			data:     []string{"second", "third"},
			nodeData: "second",
			newData:  "first",
			want:     []string{"first", "second", "third"},
		},
		{
			name:     "Append before middle element",
			data:     []string{"first", "third", "fourth"},
			nodeData: "third",
			newData:  "second",
			want:     []string{"first", "second", "third", "fourth"},
		},
		{
			name:     "Append before last element",
			data:     []string{"first", "second"},
			nodeData: "second",
			newData:  "middle",
			want:     []string{"first", "middle", "second"},
		},
		{
			name:     "Append before non-existent element",
			data:     []string{"first", "second"},
			nodeData: "third",
			newData:  "fourth",
			want:     []string{"first", "second"},
		},
		{
			name:     "Append empty string before element",
			data:     []string{"first", "second"},
			nodeData: "second",
			newData:  "",
			want:     []string{"first", "", "second"},
		},
		{
			name:     "Append before empty string element",
			data:     []string{"first", "", "third"},
			nodeData: "",
			newData:  "second",
			want:     []string{"first", "second", "", "third"},
		},
	}

	for _, test := range tests {

		t.Run(
			test.name,
			func(t *testing.T) {

				l := getInitialList(test)

				l.AppendBefore(test.nodeData, test.newData)

				testValidator(l, t, test)

			},
		)

	}

}
func TestDeleteHead(t *testing.T) {
	tests := []TestCase{
		{
			name: "Delete head from list with multiple elements",
			data: []string{"first", "second", "third"},
			want: []string{"second", "third"},
		},
		{
			name: "Delete head from list with single element",
			data: []string{"only"},
			want: []string{},
		},
		{
			name: "Delete head from empty list",
			data: []string{},
			want: []string{},
		},
		{
			name: "Delete head with empty string as first element",
			data: []string{"", "second", "third"},
			want: []string{"second", "third"},
		},
	}

	for _, test := range tests {

		t.Run(
			test.name,
			func(t *testing.T) {

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
			data: []string{"first", "second", "third"},
			want: []string{"first", "second"},
		},
		{
			name: "Delete tail from list with single element",
			data: []string{"only"},
			want: []string{},
		},
		{
			name: "Delete tail from empty list",
			data: []string{},
			want: []string{},
		},
		{
			name: "Delete tail with empty string as last element",
			data: []string{"first", "second", ""},
			want: []string{"first", "second"},
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

func TestDeleteTailMultipleTimes(t *testing.T) {
	tests := []struct {
		name           string
		data           []string
		deletionsCount int
		want           []string
	}{
		{
			name:           "Delete tail twice from list with three elements",
			data:           []string{"first", "second", "third"},
			deletionsCount: 2,
			want:           []string{"first"},
		},
		{
			name:           "Delete tail three times from list with three elements",
			data:           []string{"first", "second", "third"},
			deletionsCount: 3,
			want:           []string{},
		},
		{
			name:           "Delete tail more times than elements in the list",
			data:           []string{"first", "second"},
			deletionsCount: 3,
			want:           []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			for i := 0; i < test.deletionsCount; i++ {
				l.DeleteTail()
			}
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestDeleteTailAndAppend(t *testing.T) {
	tests := []struct {
		name       string
		data       []string
		operations []struct {
			op   string
			data string
		}
		want []string
	}{
		{
			name: "Delete tail and append new element",
			data: []string{"first", "second", "third"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "delete", data: ""},
				{op: "append", data: "fourth"},
			},
			want: []string{"first", "second", "fourth"},
		},
		{
			name: "Delete all elements and append new element",
			data: []string{"first", "second"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "delete", data: ""},
				{op: "delete", data: ""},
				{op: "append", data: "new"},
			},
			want: []string{"new"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			for _, op := range test.operations {
				if op.op == "delete" {
					l.DeleteTail()
				} else if op.op == "append" {
					l.Append(op.data)
				}
			}
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}
func TestDeleteOneBefore(t *testing.T) {
	tests := []struct {
		name     string
		data     []string
		nodeData string
		want     []string
	}{
		{
			name:     "Delete node before middle element",
			data:     []string{"first", "second", "third", "fourth"},
			nodeData: "third",
			want:     []string{"first", "third", "fourth"},
		},
		{
			name:     "Delete node before last element",
			data:     []string{"first", "second", "third"},
			nodeData: "third",
			want:     []string{"first", "third"},
		},
		{
			name:     "Delete node before first element (no change)",
			data:     []string{"first", "second", "third"},
			nodeData: "first",
			want:     []string{"first", "second", "third"},
		},
		{
			name:     "Delete in list with two elements",
			data:     []string{"first", "second"},
			nodeData: "second",
			want:     []string{"second"},
		},
		{
			name:     "Delete with non-existent node data",
			data:     []string{"first", "second", "third"},
			nodeData: "fourth",
			want:     []string{"first", "second", "third"},
		},
		{
			name:     "Delete in empty list",
			data:     []string{},
			nodeData: "first",
			want:     []string{},
		},
		{
			name:     "Delete before empty string node",
			data:     []string{"first", "", "third"},
			nodeData: "",
			want:     []string{"", "third"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			l.DeleteOneBefore(test.nodeData)
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestDeleteOneBeforeMultipleTimes(t *testing.T) {
	tests := []struct {
		name      string
		data      []string
		deletions []string
		want      []string
	}{
		{
			name:      "Delete before multiple elements",
			data:      []string{"first", "second", "third", "fourth", "fifth"},
			deletions: []string{"third", "fifth"},
			want:      []string{"first", "third", "fifth"},
		},
		{
			name:      "Delete before same element multiple times",
			data:      []string{"first", "second", "third", "fourth", "fifth"},
			deletions: []string{"fifth", "fifth", "fifth"},
			want:      []string{"first", "fifth"},
		},
		{
			name:      "Delete before alternating elements",
			data:      []string{"first", "second", "third", "fourth", "fifth", "sixth"},
			deletions: []string{"second", "fourth", "sixth"},
			want:      []string{"second", "fourth", "sixth"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			for _, nodeData := range test.deletions {
				l.DeleteOneBefore(nodeData)
			}
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestDeleteOneBeforeAndOtherOperations(t *testing.T) {
	tests := []struct {
		name       string
		data       []string
		operations []struct {
			op   string
			data string
		}
		want []string
	}{
		{
			name: "Delete before and append",
			data: []string{"first", "second", "third", "fourth"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "deleteBefore", data: "third"},
				{op: "append", data: "fifth"},
			},
			want: []string{"first", "third", "fourth", "fifth"},
		},
		{
			name: "Delete before and prepend",
			data: []string{"first", "second", "third", "fourth"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "deleteBefore", data: "fourth"},
				{op: "prepend", data: "zero"},
			},
			want: []string{"zero", "first", "second", "fourth"},
		},
		{
			name: "Delete before, append, and delete tail",
			data: []string{"first", "second", "third", "fourth"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "deleteBefore", data: "third"},
				{op: "append", data: "fifth"},
				{op: "deleteTail", data: ""},
			},
			want: []string{"first", "third", "fourth"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			for _, op := range test.operations {
				switch op.op {
				case "deleteBefore":
					l.DeleteOneBefore(op.data)
				case "append":
					l.Append(op.data)
				case "prepend":
					l.Prepend(op.data)
				case "deleteTail":
					l.DeleteTail()
				}
			}
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestDeleteOneAfter(t *testing.T) {
	tests := []TestCase{
		{
			name:     "Delete node after middle element",
			data:     []string{"first", "second", "third", "fourth"},
			nodeData: "second",
			want:     []string{"first", "second", "fourth"},
		},
		{
			name:     "Delete node after first element",
			data:     []string{"first", "second", "third"},
			nodeData: "first",
			want:     []string{"first", "third"},
		},
		{
			name:     "Delete node after last element (no change)",
			data:     []string{"first", "second", "third"},
			nodeData: "third",
			want:     []string{"first", "second", "third"},
		},
		{
			name:     "Delete in list with two elements",
			data:     []string{"first", "second"},
			nodeData: "first",
			want:     []string{"first"},
		},
		{
			name:     "Delete with non-existent node data",
			data:     []string{"first", "second", "third"},
			nodeData: "fourth",
			want:     []string{"first", "second", "third"},
		},
		{
			name:     "Delete in empty list",
			data:     []string{},
			nodeData: "first",
			want:     []string{},
		},
		{
			name:     "Delete after empty string node",
			data:     []string{"first", "", "third"},
			nodeData: "",
			want:     []string{"first", ""},
		},
	}

	for _, test := range tests {
		t.Run(
			test.name,
			func(t *testing.T) {

				l := getInitialList(TestCase{data: test.data})

				l.DeleteOneAfter(test.nodeData)

				testValidator(l, t, TestCase{want: test.want})

			})

	}

}

func TestDeleteOneAfterMultipleTimes(t *testing.T) {
	tests := []struct {
		name      string
		data      []string
		deletions []string
		want      []string
	}{
		{
			name:      "Delete after multiple elements",
			data:      []string{"first", "second", "third", "fourth", "fifth"},
			deletions: []string{"first", "third"},
			want:      []string{"first", "third", "fifth"},
		},
		{
			name:      "Delete after same element multiple times",
			data:      []string{"first", "second", "third", "fourth", "fifth"},
			deletions: []string{"first", "first", "first"},
			want:      []string{"first", "fifth"},
		},
		{
			name:      "Delete after alternating elements",
			data:      []string{"first", "second", "third", "fourth", "fifth", "sixth"},
			deletions: []string{"first", "third", "fifth"},
			want:      []string{"first", "second", "third", "fourth", "fifth"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			for _, nodeData := range test.deletions {
				l.DeleteOneAfter(nodeData)
			}
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestDeleteOneAfterAndOtherOperations(t *testing.T) {
	tests := []struct {
		name       string
		data       []string
		operations []struct {
			op   string
			data string
		}
		want []string
	}{
		{
			name: "Delete after and append",
			data: []string{"first", "second", "third", "fourth"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "deleteAfter", data: "second"},
				{op: "append", data: "fifth"},
			},
			want: []string{"first", "second", "fourth", "fifth"},
		},
		{
			name: "Delete after and prepend",
			data: []string{"first", "second", "third", "fourth"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "deleteAfter", data: "first"},
				{op: "prepend", data: "zero"},
			},
			want: []string{"zero", "first", "third", "fourth"},
		},
		{
			name: "Delete after, append, and delete head",
			data: []string{"first", "second", "third", "fourth"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "deleteAfter", data: "second"},
				{op: "append", data: "fifth"},
				{op: "deleteHead", data: ""},
			},
			want: []string{"second", "fourth", "fifth"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			for _, op := range test.operations {
				switch op.op {
				case "deleteAfter":
					l.DeleteOneAfter(op.data)
				case "append":
					l.Append(op.data)
				case "prepend":
					l.Prepend(op.data)
				case "deleteHead":
					l.DeleteHead()
				}
			}
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}
func TestDeleteAllBefore(t *testing.T) {
	tests := []TestCase{
		{
			name:     "Delete all before middle element",
			data:     []string{"first", "second", "third", "fourth", "fifth"},
			nodeData: "third",
			want:     []string{"third", "fourth", "fifth"},
		},
		{
			name:     "Delete all before last element",
			data:     []string{"first", "second", "third", "fourth"},
			nodeData: "fourth",
			want:     []string{"fourth"},
		},
		{
			name:     "Delete all before first element (no change)",
			data:     []string{"first", "second", "third"},
			nodeData: "first",
			want:     []string{"first", "second", "third"},
		},
		{
			name:     "Delete with non-existent node data",
			data:     []string{"first", "second", "third"},
			nodeData: "fourth",
			want:     []string{"first", "second", "third"},
		},
		{
			name:     "Delete in empty list",
			data:     []string{},
			nodeData: "first",
			want:     []string{},
		},
		{
			name:     "Delete all before empty string node",
			data:     []string{"first", "second", "", "third"},
			nodeData: "",
			want:     []string{"", "third"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			l.DeleteAllBefore(test.nodeData)
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestDeleteAllBeforeMultipleTimes(t *testing.T) {
	tests := []struct {
		name      string
		data      []string
		deletions []string
		want      []string
	}{
		{
			name:      "Delete all before multiple elements",
			data:      []string{"first", "second", "third", "fourth", "fifth"},
			deletions: []string{"third", "fifth"},
			want:      []string{"fifth"},
		},
		{
			name:      "Delete all before same element multiple times",
			data:      []string{"first", "second", "third", "fourth", "fifth"},
			deletions: []string{"third", "third", "third"},
			want:      []string{"third", "fourth", "fifth"},
		},
		{
			name:      "Delete all before with increasing indices",
			data:      []string{"first", "second", "third", "fourth", "fifth", "sixth"},
			deletions: []string{"second", "fourth", "sixth"},
			want:      []string{"sixth"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			for _, nodeData := range test.deletions {
				l.DeleteAllBefore(nodeData)
			}
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestDeleteAllBeforeAndOtherOperations(t *testing.T) {
	tests := []struct {
		name       string
		data       []string
		operations []struct {
			op   string
			data string
		}
		want []string
	}{
		{
			name: "Delete all before and append",
			data: []string{"first", "second", "third", "fourth"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "deleteAllBefore", data: "third"},
				{op: "append", data: "fifth"},
			},
			want: []string{"third", "fourth", "fifth"},
		},
		{
			name: "Delete all before and prepend",
			data: []string{"first", "second", "third", "fourth"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "deleteAllBefore", data: "fourth"},
				{op: "prepend", data: "zero"},
			},
			want: []string{"zero", "fourth"},
		},
		{
			name: "Delete all before, append, and delete tail",
			data: []string{"first", "second", "third", "fourth", "fifth"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "deleteAllBefore", data: "third"},
				{op: "append", data: "sixth"},
				{op: "deleteTail", data: ""},
			},
			want: []string{"third", "fourth", "fifth"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			for _, op := range test.operations {
				switch op.op {
				case "deleteAllBefore":
					l.DeleteAllBefore(op.data)
				case "append":
					l.Append(op.data)
				case "prepend":
					l.Prepend(op.data)
				case "deleteTail":
					l.DeleteTail()
				}
			}
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}
func TestDeleteAllAfter(t *testing.T) {
	tests := []struct {
		name     string
		data     []string
		nodeData string
		want     []string
	}{
		{
			name:     "Delete all after middle element",
			data:     []string{"first", "second", "third", "fourth", "fifth"},
			nodeData: "third",
			want:     []string{"first", "second", "third"},
		},
		{
			name:     "Delete all after first element",
			data:     []string{"first", "second", "third", "fourth"},
			nodeData: "first",
			want:     []string{"first"},
		},
		{
			name:     "Delete all after last element (no change)",
			data:     []string{"first", "second", "third"},
			nodeData: "third",
			want:     []string{"first", "second", "third"},
		},
		{
			name:     "Delete with non-existent node data",
			data:     []string{"first", "second", "third"},
			nodeData: "fourth",
			want:     []string{"first", "second", "third"},
		},
		{
			name:     "Delete in empty list",
			data:     []string{},
			nodeData: "first",
			want:     []string{},
		},
		{
			name:     "Delete all after empty string node",
			data:     []string{"first", "", "second", "third"},
			nodeData: "",
			want:     []string{"first", ""},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			l.DeleteAllAfter(test.nodeData)
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestDeleteAllAfterMultipleTimes(t *testing.T) {
	tests := []struct {
		name      string
		data      []string
		deletions []string
		want      []string
	}{
		{
			name:      "Delete all after multiple elements",
			data:      []string{"first", "second", "third", "fourth", "fifth"},
			deletions: []string{"second", "fourth"},
			want:      []string{"first", "second"},
		},
		{
			name:      "Delete all after same element multiple times",
			data:      []string{"first", "second", "third", "fourth", "fifth"},
			deletions: []string{"second", "second", "second"},
			want:      []string{"first", "second"},
		},
		{
			name:      "Delete all after with decreasing indices",
			data:      []string{"first", "second", "third", "fourth", "fifth", "sixth"},
			deletions: []string{"fifth", "third", "first"},
			want:      []string{"first"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			for _, nodeData := range test.deletions {
				l.DeleteAllAfter(nodeData)
			}
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestDeleteAllAfterAndOtherOperations(t *testing.T) {
	tests := []struct {
		name       string
		data       []string
		operations []struct {
			op   string
			data string
		}
		want []string
	}{
		{
			name: "Delete all after and append",
			data: []string{"first", "second", "third", "fourth"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "deleteAllAfter", data: "second"},
				{op: "append", data: "fifth"},
			},
			want: []string{"first", "second", "fifth"},
		},
		{
			name: "Delete all after and prepend",
			data: []string{"first", "second", "third", "fourth"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "deleteAllAfter", data: "second"},
				{op: "prepend", data: "zero"},
			},
			want: []string{"zero", "first", "second"},
		},
		{
			name: "Delete all after, prepend, and delete head",
			data: []string{"first", "second", "third", "fourth", "fifth"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "deleteAllAfter", data: "third"},
				{op: "prepend", data: "zero"},
				{op: "deleteHead", data: ""},
			},
			want: []string{"first", "second", "third"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			for _, op := range test.operations {
				switch op.op {
				case "deleteAllAfter":
					l.DeleteAllAfter(op.data)
				case "append":
					l.Append(op.data)
				case "prepend":
					l.Prepend(op.data)
				case "deleteHead":
					l.DeleteHead()
				}
			}
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}
func TestReverse(t *testing.T) {
	tests := []TestCase{
		{
			name: "Reverse list with multiple elements",
			data: []string{"first", "second", "third", "fourth"},
			want: []string{"fourth", "third", "second", "first"},
		},
		{
			name: "Reverse list with two elements",
			data: []string{"first", "second"},
			want: []string{"second", "first"},
		},
		{
			name: "Reverse list with one element",
			data: []string{"only"},
			want: []string{"only"},
		},
		{
			name: "Reverse empty list",
			data: []string{},
			want: []string{},
		},
		{
			name: "Reverse list with duplicate elements",
			data: []string{"first", "second", "second", "first"},
			want: []string{"first", "second", "second", "first"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			l.Reverse()
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestReverseMultipleTimes(t *testing.T) {
	tests := []struct {
		name         string
		data         []string
		reverseTimes int
		want         []string
	}{
		{
			name:         "Reverse even number of times",
			data:         []string{"first", "second", "third"},
			reverseTimes: 2,
			want:         []string{"first", "second", "third"},
		},
		{
			name:         "Reverse odd number of times",
			data:         []string{"first", "second", "third"},
			reverseTimes: 3,
			want:         []string{"third", "second", "first"},
		},
		{
			name:         "Reverse many times",
			data:         []string{"first", "second", "third"},
			reverseTimes: 10,
			want:         []string{"first", "second", "third"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			for i := 0; i < test.reverseTimes; i++ {
				l.Reverse()
			}
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}

func TestReverseAndOtherOperations(t *testing.T) {
	tests := []struct {
		name       string
		data       []string
		operations []struct {
			op   string
			data string
		}
		want []string
	}{
		{
			name: "Reverse and append",
			data: []string{"first", "second", "third"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "reverse", data: ""},
				{op: "append", data: "fourth"},
			},
			want: []string{"third", "second", "first", "fourth"},
		},
		{
			name: "Reverse and prepend",
			data: []string{"first", "second", "third"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "reverse", data: ""},
				{op: "prepend", data: "zero"},
			},
			want: []string{"zero", "third", "second", "first"},
		},
		{
			name: "Reverse, append, and reverse again",
			data: []string{"first", "second", "third"},
			operations: []struct {
				op   string
				data string
			}{
				{op: "reverse", data: ""},
				{op: "append", data: "fourth"},
				{op: "reverse", data: ""},
			},
			want: []string{"fourth", "first", "second", "third"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := getInitialList(TestCase{data: test.data})
			for _, op := range test.operations {
				switch op.op {
				case "reverse":
					l.Reverse()
				case "append":
					l.Append(op.data)
				case "prepend":
					l.Prepend(op.data)
				}
			}
			testValidator(l, t, TestCase{want: test.want})
		})
	}
}
