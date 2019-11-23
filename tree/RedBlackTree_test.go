package tree

import "testing"

func TestRedBlackTree_Insert(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(1)
	tree.Insert(10)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(2)
}
