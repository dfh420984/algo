package link

import "fmt"

type Node struct {
	No   int
	next *Node
}

//AddNodeCircle 创建一个环形链表
func AddNodeCircle(n int) *Node {
	if n <= 0 {
		return nil
	}
	head := &Node{No: 0}
	tail := head
	for i := 1; i <= n; i++ {
		node := &Node{
			No: i,
		}
		if i == 1 {
			head = node
			tail = node
			tail.next = head
		} else {
			tail.next = node
			tail = node
			tail.next = head
		}
	}
	return head
}

//ShowNodeCircle 打印环形链表
func ShowNodeCircle() {
	head := AddNodeCircle(10)
	if head == nil {
		return
	}
	cur := head
	for {
		fmt.Printf("%d->", cur.No)
		cur = cur.next
		if cur == head {
			fmt.Printf("%d", cur.No)
			break
		}
	}
	fmt.Println()
}

//PlayGame first 环形链表的环节点 startNo 起始节点，m 跳跃次数
func PlayGame(first *Node, startNo int, m int) {
	if first.next == nil || m <= 0 {
		return
	}
	cur, tail := first, first
	//找到尾节点
	for tail.next != first {
		tail = tail.next
	}
	//找到起始节点
	for i := 0; i < startNo; i++ {
		cur = cur.next
		tail = tail.next
	}
	for {
		//跳跃
		for i := 0; i < m; i++ {
			cur = cur.next
			tail = tail.next
		}
		//移除节点
		fmt.Printf("%d 被移除\n", cur.No)
		cur = cur.next
		tail.next = cur
		if cur == tail {
			break
		}
	}

}
