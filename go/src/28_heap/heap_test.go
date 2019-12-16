package _28_heap

import "testing"

func TestHeap(t *testing.T) {
	heap := NewHeap(6)
	heap.Insert(3)
	heap.Insert(6)
	heap.Insert(30)
	heap.Insert(20)
	heap.Insert(10)
	heap.Insert(60)
	t.Log(heap.count)
	t.Log(heap.a)
	heap.RemoveMax()
	t.Log(heap.count)
	t.Log(heap.a)
}
