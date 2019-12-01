package tree

type TrieTree struct {
	head *TrieNode
}

type TrieNode struct {
	data     byte
	children []*TrieNode
	isEnd    bool
}

func NewTrieNode(data byte, isEnd bool) *TrieNode {
	return &TrieNode{data, make([]*TrieNode, 26), isEnd}
}

func NewTrieTree() *TrieTree {
	return &TrieTree{NewTrieNode('/', true)}
}

func (t *TrieTree) insert(value string) {
	tmp := t.head
	for _, ch := range value {
		index := ch - 'a'
		n := tmp.children[index]
		if n == nil {
			tmp.isEnd = false
			tmp.children[index] = NewTrieNode(byte(ch), true)
			tmp = tmp.children[index]
		} else {
			tmp = n
		}
	}
}

func (t *TrieTree) find(value string) bool {
	tmp := t.head
	for _, ch := range value {
		index := ch - 'a'
		n := tmp.children[index]
		if n == nil {
			return false
		}
		tmp = tmp.children[index]
	}
	return tmp.isEnd
}
