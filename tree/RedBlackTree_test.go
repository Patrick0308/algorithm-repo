package tree

import (
	"fmt"
	"testing"
)

func TestRedBlackTree_Insert(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(1)
	tree.Insert(10)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(2)
}

func TestRedBlackTree_Search(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Insert(1)
	tree.Insert(10)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(2)

	node := tree.Search(10)
	fmt.Println(node == nil || node.value != 10)
}
