package _8_stack

import "fmt"

type Node struct {
	next *Node
	val  interface{}
}

type LinkedListStack struct {
	topNode *Node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		nil,
	}
}

func (this *LinkedListStack) IsEmpty() bool {
	if this.topNode == nil {
		return true
	} else {
		return false
	}
}

func (this *LinkedListStack) Push(v interface{}) {
	this.topNode = &Node{next: this.topNode, val: v}
}

func (this *LinkedListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	val := this.topNode.val
	this.topNode = this.topNode.next
	return val
}

func (this *LinkedListStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	val := this.topNode.val
	return val
}

func (this *LinkedListStack) Flush() {
	this.topNode = nil
}

func (this *LinkedListStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := this.topNode
		for cur != nil {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
