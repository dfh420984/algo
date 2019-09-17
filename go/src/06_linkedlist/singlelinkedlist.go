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
func (this *LinkedList) SearchNode(index uint) *LinkNode {
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

//链表反转
func (this *LinkedList) ReverseLinkList(head *LinkNode) *LinkNode {
	var pre *LinkNode = nil
	var next *LinkNode = nil
	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}

//判断链表是否是回文
func (this *LinkedList) IsLinkPalindrome() bool {
	//定义快慢两个节点指针
	var slow *LinkNode = this.head.next
	var fast *LinkNode = this.head.next
	var mid *LinkNode = nil  //中间点
	var pre *LinkNode = nil  //前一个节点
	var next *LinkNode = nil //后一个节点

	//找到中间点
	for fast != nil && fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		mid = slow.next
	}

	//判断链表是否是偶数
	if int(this.length)%2 == 0 {
		mid = mid.next
	}

	//反转链表
	for mid != nil {
		next = mid.next
		mid.next = pre
		pre = mid
		mid = next
	}

	//开始比较
	for slow != nil && pre != nil {
		if slow.data != pre.data {
			return false
		}
		slow = slow.next
		pre = pre.next
	}
	return true
}

//合并两个有序链表
func (this *LinkedList) MergeLinkList(l0, l1, l2 *LinkedList) *LinkedList {
	if l1 == nil && l1.head == nil && l1.head.next == nil {
		return l2
	}
	if l2 == nil && l2.head == nil && l2.head.next == nil {
		return l1
	}
	//初始化指针，并指向表头
	cur := l0.head
	cur1 := l1.head.next
	cur2 := l2.head.next
	//开始循环遍历两个链表数据，数据小的插入在新链表后面
	for cur1 != nil && cur2 != nil {
		if cur1.data.(int) < cur2.data.(int) {
			cur.next = cur1
			cur1 = cur1.next
		} else {
			cur.next = cur2
			cur2 = cur2.next
		}
		cur = cur.next
	}
	//循环完毕后，链表剩余的节点插入
	if cur1 != nil {
		cur.next = cur1
	}
	if cur2 != nil {
		cur.next = cur2
	}
	return l0
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

func (this *LinkedList) PrintNode(head *LinkNode) {
	var format string = ""
	for head != nil {
		format += fmt.Sprintf("|%+v", head.data)
		head = head.next
		if head != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
