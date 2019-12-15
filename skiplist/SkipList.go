package skiplist

import (
	"math"
	"math/rand"
)

const MaxLevel = 16

type skipListNode struct {
	v interface{}

	score int

	level int

	forwards []*skipListNode
}

func newSkipListNode(v interface{}, score, level int) *skipListNode {
	return &skipListNode{
		v:        v,
		score:    score,
		level:    level,
		forwards: make([]*skipListNode, level, level),
	}
}

type SkipList struct {
	head *skipListNode

	level int

	length int
}

func NewSkipList() *SkipList {
	return &SkipList{
		head:   newSkipListNode(nil, math.MinInt32, MaxLevel),
		level:  0,
		length: 0,
	}
}

func (l *SkipList) Insert(v interface{}, score int) bool {
	if nil == v {
		return false
	}
	cur := l.head
	updateNode := [MaxLevel]*skipListNode{}
	for i := MaxLevel - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i] == v {
				return false
			}
			if cur.forwards[i].score > score {
				updateNode[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if nil == cur.forwards[i] {
			updateNode[i] = cur
		}
	}
	level := randLevel()
	newNode := newSkipListNode(v, score, level)

	for i := 0; i < level; i++ {
		next := updateNode[i].forwards[i]
		updateNode[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	if l.level < level {
		l.level = level
	}
	l.length++
	return true
}

func randLevel() int {
	level := 1
	for i := 1; i < MaxLevel; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}
	return level
}
