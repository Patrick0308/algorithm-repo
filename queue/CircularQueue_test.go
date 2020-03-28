package queue

import (
	"fmt"
	"testing"
)

func TestCircularQueue_EnQueue(t *testing.T) {
	cq := NewCircularQueue(10)
	cq.EnQueue(1)
	cq.EnQueue(1)
	fmt.Println(cq.DeQueue())
	fmt.Println(cq.DeQueue())
}
