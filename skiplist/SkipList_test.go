package skiplist

import "testing"

func TestSkipList_Insert(t *testing.T) {
	sl := NewSkipList()
	sl.Insert(1, 1)
	sl.Insert(2, 2)
	sl.Insert(10, 10)
	sl.Insert(5, 5)
	sl.Insert(11, 11)
}
