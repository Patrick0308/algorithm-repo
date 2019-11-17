package stack

import (
	"fmt"
	"testing"
)

func TestArrayStack_Pop(t *testing.T) {
	s := NewArrayStack()
	for i := 0; i < 10; i++ {
		s.Push(i + 1)
	}
	for i := 0; i < 10; i++ {
		fmt.Print(s.Pop())
		fmt.Print("->")
	}
	fmt.Println("End")
}

func TestArrayStack_Push(t *testing.T) {
	s := NewArrayStack()
	for i := 0; i < 10; i++ {
		s.Push(i + 1)
	}
	fmt.Println(s.String())
}
