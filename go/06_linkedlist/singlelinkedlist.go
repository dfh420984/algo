package linkedlist

import "fmt"

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type LinkedList struct {
	length uint
	head   *LinkNode
}

//链表节点
func NewLinkNode(v interface{}) *LinkNode {
	return &LinkNode{
		data: v,
		next: nil,
	}
}

//初始化链表,并生成一个哨兵头节点
func NewLinkedList() *LinkedList {
	return &LinkedList{
		length: 0,
		head:   NewLinkNode(0),
	}
}

//在节点后面插入节点
func (this *LinkedList) InsertAfter(oriNode *LinkNode, v interface{}) bool {
	newNode := NewLinkNode(v)
	return this.InsertAfterNode(oriNode, newNode)
}

func (this *LinkedList) InsertAfterNode(oriNode *LinkNode, newNode *LinkNode) bool {
	if oriNode == nil {
		return false
	}
	//更换节点指向
	newNode.next = oriNode.next
	oriNode.next = newNode
	this.length++
	return true
}

//在节点前面插入节点
func (this *LinkedList) InsertBefore(oriNode *LinkNode, v interface{}) bool {
	newNode := NewLinkNode(v)
	return this.InsertAfterNode(oriNode, newNode)
}

func (this *LinkedList) InsertBeforeNode(oriNode *LinkNode, newNode *LinkNode) bool {
	if oriNode == nil || this.head == oriNode {
		return false
	}
	//找到前一个节点
	preNode := this.getPreNode(oriNode)
	newNode.next = preNode.next
	preNode.next = newNode
	this.length++
	return true
}

//获取某个节点前面的前一个节点
func (this *LinkedList) getPreNode(oriNode *LinkNode) *LinkNode {
	cur := this.head
	for cur != nil {
		if cur == oriNode {
			break
		}
		cur = cur.next
	}
	if cur == nil { //没有找到节点
		return nil
	}
	return cur
}

//在链表头部插入节点
func (this *LinkedList) InserToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

//在链表尾部插入节点
func (this *LinkedList) InsertTotail(v interface{}) bool {
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

//根据下标查找节点
func (this *LinkedList) searchNode(index uint) *LinkNode {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//删除传入节点
func (this *LinkedList) DeleteNode(p *LinkNode) bool {
	if p == nil {
		return false
	}
	cur := this.head.next
	pre := this.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	pre.next = p.next
	p = nil
	this.length--
	return true
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("|%+v", cur.data)
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
