package arraystack

import "fmt"

type IArrayStack interface {
	Push(data any)
	Pop() any
	Peek() any
	Length()
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
	data []any
	top  int
}

func NewArrayStack(capacity int) *ArrayStack {

	if capacity < 0 {
		capacity = 0
	}

	return &ArrayStack{
		data: make([]any, 0, capacity),
		top:  -1,
	}

}

func (s *ArrayStack) Push(data any) {

	s.data = append(s.data, data)

	s.top++

}

func (s *ArrayStack) Pop() any {

	if s.top == -1 {
		return nil
	}

	data := s.data[s.top]

	s.data = s.data[:s.top]

	s.top--

	return data

}

func (s *ArrayStack) Peek() any {

	if s.top == -1 {
		return nil
	}

	return s.data[s.top]

}

func (s *ArrayStack) IsEmpty() bool {
	return s.top == -1
}

func (s *ArrayStack) Length() {
	fmt.Println("\nLenght: ", s.top+1)
}

func (s *ArrayStack) PrintStack() {

	fmt.Printf("Stack: %+v\n\n", s)

	for i := s.top; i >= 0; i-- {
		fmt.Printf("| %v |\n", s.data[i])
	}

}

var _ IArrayStack = (*ArrayStack)(nil)
