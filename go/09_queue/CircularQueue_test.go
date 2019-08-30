package _9_queue

import "testing"

func TestCircularQueue(t *testing.T) {
	q := NewCircularQueue(4)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	t.Log(q)
	t.Log(q.DeQueue())
	t.Log(q.DeQueue())
	t.Log(q.DeQueue())
	t.Log(q.DeQueue())
	t.Log(q)
}
