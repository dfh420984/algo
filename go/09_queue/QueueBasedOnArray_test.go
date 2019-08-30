package _9_queue

import (
	"fmt"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	q := NewArrayQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	str := q.String()
	fmt.Println(str)
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	str = q.String()
	fmt.Println(str)
}
