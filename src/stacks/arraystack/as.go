package arraystack

import "fmt"

type IArrayStack interface {
	Push(data interface{})
	Pop() interface{}
	Peek() interface{}
	Length() int
	IsEmpty() bool
	PrintStack()
}

func GetChoices() []string {

	return []string{
		"Push",
		"Pop",
		"Peek",
		"Length",
		"Is empty",
		"Print stack",
	}

}

type ArrayStack struct {
	data   []interface{}
	top    int
	length int
}

func NewArrayStack(capacity int) *ArrayStack {

	if capacity < 0 {
		capacity = 0
	}

	a := ArrayStack{
		data:   make([]interface{}, capacity),
		top:    -1,
		length: 0,
	}

	a.data = a.data[:0]

	return &a

}

func (s *ArrayStack) Push(data interface{}) {

	s.top++

	s.data = append(s.data, data)

	s.length++

}

func (s *ArrayStack) Pop() interface{} {

	if s.top == -1 {
		return nil
	}

	data := s.data[s.top]

	s.top--

	s.length--

	s.data = s.data[:len(s.data)-1]

	return data

}

func (s *ArrayStack) Peek() interface{} {

	if s.top == -1 {
		return nil
	}

	return s.data[s.top]

}

func (s *ArrayStack) IsEmpty() bool {
	return s.top == -1
}

func (s *ArrayStack) Length() int {
	return s.length
}

func (s *ArrayStack) PrintStack() {

	fmt.Printf("Stack: %+v\n\n", s)

	for i := s.top; i >= 0; i-- {
		fmt.Printf("| %v |\n", s.data[i])
	}

}

var _ IArrayStack = (*ArrayStack)(nil)
