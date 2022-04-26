package link

import (
	"testing"
)

func TestShowLink(t *testing.T) {
	l := &LinkList{}
	l.AddNodeToTail(1)
	l.AddNodeToTail(2)
	l.AddNodeToTail(3)
	l.ShowLink(l.head)
}

func TestReverseLink(t *testing.T) {
	l := &LinkList{}
	l.AddNodeToTail(1)
	l.AddNodeToTail(2)
	l.AddNodeToTail(3)
	l.AddNodeToTail(2)
	//l.ShowLink(l.head)
	//node := l.ReverseLink(l.head)
	//l.ShowLink(node)
	flag := l.IsPalindromList(l.head)
	t.Logf("%v", flag)
}

func TestMergeSortLinkList(t *testing.T) {
	ll0 := &LinkList{}
	ll0.AddNodeToTail(0)
	ll1 := &LinkList{}
	ll1.AddNodeToTail(0)
	ll1.AddNodeToTail(1)
	ll1.AddNodeToTail(3)
	ll1.AddNodeToTail(5)
	ll2 := &LinkList{}
	ll2.AddNodeToTail(0)
	ll2.AddNodeToTail(2)
	ll2.AddNodeToTail(4)
	ll2.AddNodeToTail(6)
	ml := MergeSortLinkList(ll0.head, ll1.head, ll2.head)
	l := &LinkList{}
	l.ShowLink(ml)
}
