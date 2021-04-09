package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue(3)
	q.enqueue("1")
	q.enqueue("2")
	q.enqueue("3")
	println(q.enqueue("4"))
	println(q.dequeue())
	println(q.dequeue())
	println(q.dequeue())
}

type Queue struct {
	item []string
	len  int
	cap  int
}

func NewQueue(cap int) *Queue {
	var item []string
	return &Queue{
		item: item,
		len:  0,
		cap:  cap,
	}
}

func (q *Queue) enqueue(s string) bool {
	if q.len >= q.cap {
		return false
	}
	q.item = append(q.item, s)
	q.len++
	return true
}

func (q *Queue) dequeue() string {
	if q.len == 0 {
		return ""
	}
	tmp := q.item[q.cap-q.len]
	q.len--
	return tmp
}

