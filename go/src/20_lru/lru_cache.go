package _20_lru

const (
	hostbit = uint64(^uint(0)) == ^uint64(0)
	LENGTH  = 100
)

type lruNode struct {
	prev  *lruNode //前一节点指针
	next  *lruNode //后一节点指针
	key   int      //lru节点key
	value int      //lru节点value
	hnext *lruNode //拉链
}

type lruCache struct {
	node     []lruNode //hash list
	head     *lruNode  // lru head node
	tail     *lruNode  // lru tail node
	capacity int
	used     int
}

//实例lru
func NewLruCache(capacity int) *lruCache {
	return &lruCache{
		node:     make([]lruNode, LENGTH),
		head:     nil,
		tail:     nil,
		capacity: capacity,
		used:     0,
	}
}

func (this *lruCache) Get(key int) int {
	if this.tail == nil {
		return -1
	}

	if tmp := this.searchNode(key); tmp != nil {
		this.moveToTail(tmp)
		return tmp.value
	}
	return -1
}

func (this *lruCache) Put(key int, value int) {
	// 1. 首次插入数据
	// 2. 插入数据不在 LRU 中
	// 3. 插入数据在 LRU 中
	// 4. 插入数据不在 LRU 中, 并且 LRU 已满

	if tmp := this.searchNode(key); tmp != nil {
		tmp.value = value
		this.moveToTail(tmp)
		return
	}
	this.addNode(key, value)

	if this.used > this.capacity {
		this.delNode()
	}
}

//插入数据到lru中
func (lr *lruCache) addNode(key int, value int) {
	newNode := &lruNode{
		key:   key,
		value: value,
	}

	//将节点串联在拉链中
	tmp := &lr.node[hash(key)]
	newNode.hnext = tmp.hnext
	tmp.hnext = newNode
	lr.used++

	//处理双向链表
	if lr.tail == nil { //首次插入数据
		lr.head, lr.tail = newNode, newNode
		return
	}

	lr.tail.next = newNode
	newNode.prev = lr.tail
	lr.tail = newNode

}

//删除节点
func (lr *lruCache) delNode() {
	if lr.head == nil {
		return
	}

	prev := &lr.node[hash(lr.head.key)]
	tmp := prev.hnext

	for tmp != nil && tmp.key != lr.head.key {
		prev = tmp
		tmp = tmp.hnext
	}
	if tmp == nil {
		return
	}
	prev.hnext = tmp.hnext
	lr.head = lr.head.next
	lr.head.prev = nil
	lr.used--
}

//查找节点
func (lr *lruCache) searchNode(key int) *lruNode {
	if lr.tail == nil {
		return nil
	}

	tmp := lr.node[hash(key)].hnext

	for tmp != nil {
		if tmp.key == key {
			return tmp
		}
		tmp = tmp.hnext
	}
	return nil
}

//移到尾部
func (this *lruCache) moveToTail(node *lruNode) {
	if this.tail == node {
		return
	}
	if this.head == node {
		this.head = node.next
		this.head.prev = nil
	} else {
		node.next.prev = node.prev
		node.prev.next = node.next
	}

	node.next = nil
	this.tail.next = node
	node.prev = this.tail

	this.tail = node
}

func hash(key int) int {
	if hostbit {
		return (key ^ (key >> 32)) & (LENGTH - 1)
	}
	return (key ^ (key >> 16)) & (LENGTH - 1)
}
