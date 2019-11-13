package queue

import "fmt"

type CircularQueue struct {
	values []interface{}
	capacity int
	head int
	tail int
}

func NewCircularQueue(n int)  *CircularQueue{
	return  &CircularQueue{make([] interface{}, n), n, 0, 0}
}

func (q *CircularQueue) IsEmpty() bool {
	return q.head == q.tail
}

func (q *CircularQueue) IsFull() bool{
	return q.head == (q.tail+1)%q.capacity
}

func (q *CircularQueue) EnQueue(v interface{}) bool{
	if q.IsFull() {
		return false
	}
	q.tail = (q.tail+1)%q.capacity
	q.values[q.tail] = v
	return true
}

func (q *CircularQueue) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	value := q.values[q.head]
	q.values[q.head] = nil
	q.head = (q.head +1)%q.capacity
	return value
}

func (q *CircularQueue) String() string{
	if q.IsEmpty() {
		return ""
	}
	result := "head"
	i := q.head
	for true {
		result += fmt.Sprintf("<-%+v", q.values[i])
		i = (i + 1) % q.capacity
		if i == q.tail {
			break
		}
	}
	return ""
}

