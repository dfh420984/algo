package queue

//cirCleQueue 环形队列
type cirCleQueue struct {
	items []string //存放数据
	n     int      //队列长度
	head  int      //头下标
	tail  int      //尾下标
}

func NewcirCleQueue(n int) *cirCleQueue {
	return &cirCleQueue{
		items: make([]string, n),
		n:     n,
	}
}

//EnQueue 入队
func (c *cirCleQueue) EnQueue(item string) bool {
	//判断队列是否已满
	if (c.tail+1)%c.n == c.head { //环形队列tail指针会占用一个位置
		return false
	}
	c.items[c.tail] = item
	c.tail = (c.tail + 1) % c.n
	return true
}

//Dequeue() 出队
func (c *cirCleQueue) DeQueue() string {
	if c.head == c.tail {
		return ""
	}
	item := c.items[c.head]
	c.head = (c.head + 1) % c.n
	return item
}
