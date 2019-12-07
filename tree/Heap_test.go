package tree

import "testing"

func TestHeap_Insert(t *testing.T) {
	h := NewHeap(6)
	h.Insert(5)
	h.Insert(2)
	h.Insert(8)
	h.Insert(9)
	h.Insert(0)
}

func TestHeap_RemoveMax(t *testing.T) {
	h := NewHeap(6)
	h.Insert(5)
	h.Insert(2)
	h.Insert(8)
	h.Insert(9)
	h.Insert(0)
	h.RemoveMax()
	h.RemoveMax()
}
