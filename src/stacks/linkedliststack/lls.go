package linkedliststack

import "fmt"

type Node struct {
	data interface{}
	prev *Node
}

type Stack struct {
	top  *Node
	size int
}

func (s *Stack) Push(data interface{}) {

	newNode := &Node{data: data, prev: s.top}

	s.top = newNode

	s.size++

}

func (s *Stack) Pop() interface{} {

	if s.top == nil {
		return nil
	}

	data := s.top.data

	s.top = s.top.prev

	s.size--

	return data

}

func (s *Stack) Peek() interface{} {

	if s.top == nil {
		return nil
	}

	return s.top.data

}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) PrintStack() {

	currentNode := s.top

	for currentNode.prev != nil {
		fmt.Printf("| %v |\n", currentNode.data)
	}

}
