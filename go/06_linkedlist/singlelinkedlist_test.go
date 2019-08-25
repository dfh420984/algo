package linkedlist

import (
	"fmt"
	"testing"
)

func TestInserToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 5; i++ {
		l.InserToHead(i + 1)
	}
	l.Print()
	reverse := l.ReverseLinkList(l.head.next)
	l.PrintNode(reverse)
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
	t.Log(l.SearchNode(0))
	t.Log(l.SearchNode(2))
	t.Log(l.SearchNode(4))
	t.Log(l.SearchNode(6))
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

func TestIsLinkPalindrome(t *testing.T) {
	l := NewLinkedList()
	l.InsertTotail(1)
	l.InsertTotail(2)
	l.InsertTotail(3)
	l.InsertTotail(2)
	l.InsertTotail(1)
	l.Print()
	res := l.IsLinkPalindrome()
	fmt.Println(res)
}

func TestMergeLinkList(t *testing.T) {
	l0 := NewLinkedList()
	l1 := NewLinkedList()
	l1.InsertTotail(1)
	l1.InsertTotail(3)
	l1.InsertTotail(5)
	l1.Print()

	l2 := NewLinkedList()
	l2.InsertTotail(2)
	l2.InsertTotail(4)
	l2.InsertTotail(6)
	l2.Print()

	l0 = l0.MergeLinkList(l0, l1, l2)
	l0.Print()

}
