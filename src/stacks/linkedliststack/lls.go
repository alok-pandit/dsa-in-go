package linkedliststack

import "fmt"

type ILinkedListStack interface {
	Push(data any)
	Pop() any
	Peek() any
	IsEmpty() bool
	Size() int
	PrintStack()
}

func GetChoices() []string {

	return []string{
		"Push",
		"Pop",
		"Peek",
		"Is empty",
		"Size",
		"Print stack",
	}

}

type Node struct {
	data any
	prev *Node
}

type LinkedListStack struct {
	top  *Node
	size int
}

func (s *LinkedListStack) Push(data any) {

	s.top = &Node{data: data, prev: s.top}

	s.size++

}

func (s *LinkedListStack) Pop() any {

	if s.top == nil {
		return nil
	}

	data := s.top.data

	s.top = s.top.prev

	s.size--

	return data

}

func (s *LinkedListStack) Peek() any {

	if s.top == nil {
		return nil
	}

	return s.top.data

}

func (s *LinkedListStack) IsEmpty() bool {
	return s.top == nil
}

func (s *LinkedListStack) Size() int {
	return s.size
}

func (s *LinkedListStack) PrintStack() {

	currentNode := s.top

	for currentNode.prev != nil {
		fmt.Printf("| %v |\n", currentNode.data)
	}

}

var _ ILinkedListStack = (*LinkedListStack)(nil)
