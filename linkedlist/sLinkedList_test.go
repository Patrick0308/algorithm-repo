package linkedlist

import (
	"fmt"
	"testing"
)

func TestInsertToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertHead(i +1)
	}
	l.Print()
}

func TestInsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertTail(i + 1)
	}
	l.Print()
}

func TestLinkedList_Reverse(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertHead(i +1)
	}
	l.Print()
	l.Reverse()
	l.Print()
}

func TestLinkedList_DeleteBottomN(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertHead(i +1)
	}
	l.Print()
	l.DeleteBottomN(2)
	l.Print()
}

func TestLinkedList_FindMiddleNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertHead(i +1)
	}
	l.Print()
	fmt.Println(l.FindMiddleNode())
}