package stack

import "fmt"

/*
	Stack implement on array
*/
type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{data: make([]interface{}, 0, 32),
		top: -1}
}

func (s *ArrayStack) IsEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

func (s *ArrayStack) Push(v interface{}) {
	if s.top < 0 {
		s.top = 0
	} else {
		s.top += 1
	}
	if s.top > len(s.data)-1 {
		s.data = append(s.data, v)
	} else {
		s.data[s.top] = v
	}
}

func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.data[s.top]
	s.top--
	return v
}

func (s *ArrayStack) Flush() {
	s.top = -1
}

func (s *ArrayStack) String() string {
	if s.IsEmpty() {
		return "Empty Stack"
	}
	str := ""
	for i := s.top; i >= 0; i-- {
		str += fmt.Sprintf("<-%+v", s.data[i])
	}
	return str
}
