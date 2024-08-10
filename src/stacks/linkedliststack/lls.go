package linkedliststack

import "fmt"

type ILinkedListStack interface {
	Push(data interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmpty() bool
	Size() int
	PrintStack()
}

func GetChoices() []string {

	return []string{
		"Push",
		"Pop",
		"Peek",
		"Length",
		"Is empty",
		"Size",
		"Print stack",
	}

}

type Node struct {
	data interface{}
	prev *Node
}

type LinkedListStack struct {
	top  *Node
	size int
}

func (s *LinkedListStack) Push(data interface{}) {

	s.top = &Node{data: data, prev: s.top}

	s.size++

}

func (s *LinkedListStack) Pop() interface{} {

	if s.top == nil {
		return nil
	}

	data := s.top.data

	s.top = s.top.prev

	s.size--

	return data

}

func (s *LinkedListStack) Peek() interface{} {

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
