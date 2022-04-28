package link

import "fmt"

type LinkNode struct {
	Val  int
	next *LinkNode
}

type LinkList struct {
	length int
	head   *LinkNode
}

//AddNodeToTail 尾插法单向链表
func (l *LinkList) AddNodeToTail(val int) {
	node := &LinkNode{
		Val: val,
	}
	if l.head == nil {
		l.head = node
	} else {
		cur := l.head
		for { //找到最后一个节点
			if cur.next == nil {
				break
			}
			cur = cur.next
		}
		cur.next = node
	}
	l.length++
}

//ShowLink 打印链表
func (l *LinkList) ShowLink(head *LinkNode) {
	if head == nil {
		return
	}
	cur := head
	for cur.next != nil {
		fmt.Printf("%v->", cur.Val)
		cur = cur.next
	}
	fmt.Printf("%v \n", cur.Val)
}

//ReverseLink 反转链表
func (l *LinkList) ReverseLink(head *LinkNode) *LinkNode {
	cur := head
	if cur == nil {
		return cur
	}
	var pre *LinkNode
	for cur != nil {
		next := cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	return pre
}

//PalindromList 回文链表
func (l *LinkList) IsPalindromList(head *LinkNode) bool {
	if head == nil || head.next == nil {
		return false
	}
	//定义快慢两个指针
	slow, fast := head, head
	//找到中间节点
	for fast != nil && fast.next != nil { //快指针走两步，慢指针走一步
		slow = slow.next
		fast = fast.next.next
	}
	mid := slow
	//如果是偶数，中间节点walk一步
	if fast == nil {
		mid = mid.next
	}
	//反转后半部分链表
	revNode := l.ReverseLink(mid)
	for revNode != nil {
		if head.Val != revNode.Val {
			return false
		}
		head = head.next
		revNode = revNode.next
	}
	return true
}

//MergeSortLinkList
func MergeSortLinkList(l0, l1, l2 *LinkNode) *LinkNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l0 == nil {
		l0 = &LinkNode{} //哨兵节点，用来头插入
	}
	cur := l0
	cur1 := l1.next
	cur2 := l2.next
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			cur.next = cur1
			cur1 = cur1.next
		} else {
			cur.next = cur2
			cur2 = cur2.next
		}
		cur = cur.next
	}
	for cur1 != nil {
		cur.next = cur1
	}
	for cur2 != nil {
		cur.next = cur2
	}
	return l0.next
}
