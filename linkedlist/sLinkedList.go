package linkedlist

import (
	"fmt"
)

type Node struct{
	next *Node
	value interface{}
}

type LinkedList struct {
	head *Node
	length uint
}

func NewNode(v interface{}) *Node {
	return &Node{nil, v}
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) GetValue() interface{} {
	return n.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewNode(0), 0}
}

func (n *LinkedList) InsertAfter(p *Node, v interface{}) bool{
	if nil == p {
		return false
	}
	newNode := NewNode(v)
	newNode.next = p.next
	p.next = newNode
	n.length++
	return true
}

func (n *LinkedList) InsertBefore(p *Node, v interface{}) bool{
	if nil == p && p == n.head{
		return false
	}
	t := n.head.next
	pre := n.head
	for t != nil {
		if t == p {
			break
		}
		pre = t
		t = t.next
	}
	n.InsertAfter(pre, v)
	return true
}

func (n *LinkedList) InsertHead(v interface{}) bool{
	return n.InsertAfter(n.head, v)
}

func (n *LinkedList) InsertTail(v interface{}) bool{
	tail := n.head
	for tail.next != nil  {
		tail = tail.next
	}
	return n.InsertAfter(tail,  v)
}

func (n *LinkedList) Reverse() {
	if nil == n.head.next || nil == n.head.next.next{
		return
	}
	var pre *Node =  nil
	cur :=  n.head.next
	for cur != nil{
		t := cur.next
		cur.next = pre
		pre = cur
		cur = t
	}
	n.head.next = pre
}


func (n *LinkedList) HasCycle() bool{
	slow := n.head
	fast := n.head
	for fast != nil && fast.next.next != nil{
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

func  (n *LinkedList) DeleteBottomN(num int) {
	if num < 0 || n.head.next == nil {
		return 
	}
	fast := n.head
	for i := 0; i < num && fast != nil; i++ {
		fast = fast.next
	}
	if nil == fast {
		return
	}
	slow := n.head
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}

func (n *LinkedList) FindMiddleNode() *Node{
	if n.head.next == nil {
		return nil
	}
	slow := n.head
	fast := n.head
	for fast != nil &&  fast.next != nil{
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func (n *LinkedList) Print(){
	cur := n.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}