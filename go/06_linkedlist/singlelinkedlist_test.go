package linkedlist

import "testing"

func TestInserToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 5; i++ {
		l.InserToHead(i + 1)
	}
	l.Print()
}

func TestInsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertTotail(i + 1)
	}
	l.Print()
}

func TestSearchNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertTotail(i + 1)
	}
	t.Log(l.searchNode(0))
	t.Log(l.searchNode(2))
	t.Log(l.searchNode(4))
	t.Log(l.searchNode(6))
}

func TestDeleteNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 3; i++ {
		l.InsertTotail(i + 1)
	}
	l.Print()

	t.Log(l.DeleteNode(l.head.next))
	l.Print()

	t.Log(l.DeleteNode(l.head.next.next))
	l.Print()
}
