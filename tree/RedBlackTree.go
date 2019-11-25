package tree

type Color int

const (
	RED   Color = 0
	BLACK Color = 1
)

type TreeNode struct {
	color  Color
	value  int
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
}

func NewTreeNode(value int, color Color, parent *TreeNode) *TreeNode {
	return &TreeNode{
		color:  color,
		value:  value,
		left:   nil,
		right:  nil,
		parent: parent,
	}
}

type RedBlackTree struct {
	root *TreeNode
	size int
	leaf *TreeNode
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{
		root: nil,
		size: 0,
		leaf: NewTreeNode(0, BLACK, nil),
	}
}

func (t *RedBlackTree) Insert(value int) {
	newNode := NewTreeNode(value, RED, nil)
	newNode.left = t.leaf
	newNode.right = t.leaf
	if t.root == nil {
		newNode.color = BLACK
		t.root = newNode
		return
	}
	node := t.root
	pnode := t.root
	for {
		pnode = node
		if node.value < value {
			node = node.right
		} else {
			node = node.left
		}
		if t.leaf == node {
			break
		}
	}
	if pnode.value > value {
		pnode.left = newNode
	} else {
		pnode.right = newNode
	}
	newNode.parent = pnode
	t.fixTree(newNode)
}

func (t *RedBlackTree) fixTree(node *TreeNode) {
	n := node
	for {
		if n == nil {
			break
		}
		pnode := n.parent
		if n == t.root || pnode.color == BLACK {
			break
		}
		ppnode := pnode.parent
		var unode *TreeNode
		if ppnode.left == pnode {
			unode = ppnode.right
		} else {
			unode = ppnode.left
		}
		if unode.color == RED {
			ppnode.color = RED
			unode.color = BLACK
			pnode.color = BLACK
			n = pnode.parent
		} else if pnode == ppnode.left {
			if pnode.right == n {
				t.rotateLeft(n.parent)
			}
			t.rotateRight(ppnode)
			ppnode.color = RED
			if ppnode.parent != nil {
				ppnode.parent.color = BLACK
				n = ppnode.parent
			} else {
				n = pnode.parent
			}
		} else if pnode == ppnode.right {
			if pnode.left == n {
				t.rotateRight(n.parent)
			}
			t.rotateLeft(ppnode)
			ppnode.color = RED
			if ppnode.parent != nil {
				ppnode.parent.color = BLACK
				n = ppnode.parent
			} else {
				n = pnode.parent
			}
		}
		t.root.color = BLACK
	}
}

func (t *RedBlackTree) rotateLeft(xnode *TreeNode) {
	pnode := xnode.parent
	ynode := xnode.right

	xnode.right = ynode.left
	ynode.left = xnode
	ynode.parent = pnode
	xnode.parent = ynode
	if pnode == nil {
		t.root = ynode
		return
	}
	if pnode.left == xnode {
		pnode.left = ynode
	} else {
		pnode.right = ynode
	}
}

func (t *RedBlackTree) rotateRight(xnode *TreeNode) {
	pnode := xnode.parent
	ynode := xnode.left

	xnode.left = ynode.right
	ynode.right = xnode
	ynode.parent = pnode
	xnode.parent = ynode
	if pnode == nil {
		t.root = ynode
		return
	}
	if pnode.right == xnode {
		pnode.right = ynode
	} else {
		pnode.left = ynode
	}
}

func (t *RedBlackTree) Search(value int) *TreeNode {
	tnode := t.root
	for {
		if tnode == t.leaf || tnode.value == value {
			return tnode
		}
		if tnode.value < value {
			tnode = tnode.right
		} else {
			tnode = tnode.left
		}
	}
}
