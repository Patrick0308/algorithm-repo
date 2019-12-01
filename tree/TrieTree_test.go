package tree

import (
	"fmt"
	"testing"
)

func TestTrieTree_insert(t *testing.T) {
	tree := NewTrieTree()
	tree.insert("ab")
	tree.insert("ac")
}

func TestTrieTree_find(t *testing.T) {
	tree := NewTrieTree()
	tree.insert("ab")
	tree.insert("ac")
	tree.insert("abcd")
	tree.insert("ncced")
	fmt.Println(tree.find("ncced"))
}
