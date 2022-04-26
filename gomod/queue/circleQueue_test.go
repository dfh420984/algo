package queue

import "testing"

func TestEnQueue(t *testing.T) {
	q := NewcirCleQueue(5)
	q.EnQueue("a")
	q.EnQueue("b")
	q.EnQueue("c")
	q.EnQueue("d")
	q.EnQueue("e")
	q.EnQueue("f")
	t.Log(q)
	i := q.DeQueue()
	t.Log(i)
	i = q.DeQueue()
	t.Log(i)
	i = q.DeQueue()
	t.Log(i)
	i = q.DeQueue()
	t.Log(i)
	i = q.DeQueue()
	t.Log(i)
	i = q.DeQueue()
	t.Log(i)
}
