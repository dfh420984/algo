package _17_skiplist

import (
	"math"
	"math/rand"
)

const (
	MAX_LEVEl = 16
)

//跳表节点结构体
type skipListNode struct {
	//节点保存值
	v interface{}
	//节点排序分值
	score int
	//层高
	level int
	//每层前进指针
	forward []*skipListNode
}

//新建跳表节点
func NewSkipListNode(v interface{}, score int, level int) *skipListNode {
	return &skipListNode{
		v:       v,
		score:   score,
		level:   level,
		forward: make([]*skipListNode, level, level),
	}
}

//跳表结构体
type skipList struct {
	//跳表头节点
	head *skipListNode
	//跳表当前层数
	level int
	//跳表长度
	length int
}

//实例化跳表对象
func NewSkipList() *skipList {
	//跳表头节点,便于操作
	head := NewSkipListNode(0, math.MinInt32, MAX_LEVEl)
	return &skipList{head: head, level: 1, length: 0}
}

//获取跳表长度
func (sl *skipList) Length() int {
	return sl.length
}

//获取跳表层数
func (sl *skipList) Level() int {
	return sl.level
}

//插入节点到跳表中
func (sl *skipList) Insert(v interface{}, score int) int {
	if nil == v {
		return 1
	}
	//查找插入位置
	cur := sl.head
	//记录每层的节点路径
	update := [MAX_LEVEl]*skipListNode{}
	//从上往下逐层查找
	i := MAX_LEVEl - 1
	for ; i >= 0; i-- {
		for nil != cur.forward[i] {
			if cur.forward[i].v == v { //已经插入过
				return 2
			}

			if cur.forward[i].score > score { //待插入的分值小于已存在的索引层节点分值
				update[i] = cur
				break
			}
			cur = cur.forward[i]
		}
		if nil == cur.forward[i] {
			update[i] = cur
		}
	}

	//通过随机算法获取该节点要插入层数
	level := 1
	for i := 1; i < MAX_LEVEl; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	//创建一个新的跳表节点(待插入的节点)
	newNode := NewSkipListNode(v, score, level)

	//原有节点连接
	for i := 0; i <= level-1; i++ {
		next := update[i].forward[i]
		update[i].forward[i] = newNode
		newNode.forward[i] = next
	}

	//如果当前节点层数大于之前跳表层数，更新当前跳表层数
	if level > sl.level {
		sl.level = level
	}

	//更新跳表长度
	sl.length++

	return 0
}

//查找
func (sl *skipList) Find(v interface{}, score int) *skipListNode {
	if v == nil || sl.length == 0 {
		return nil
	}
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for nil != cur.forward[i] {
			if cur.forward[i].score == score && cur.forward[i].v == v {
				return cur.forward[i]
			} else if cur.forward[i].score > score { //跳出本层，到下层查询
				break
			}
			//本层下一个节点
			cur = cur.forward[i]
		}
	}
	return nil
}

//删除节点
func (sl *skipList) Delete(v interface{}, score int) int {
	if v == nil || sl.length == 0 {
		return 1
	}

	//查找前驱节点
	cur := sl.head
	//记录前驱路径
	update := [MAX_LEVEl]*skipListNode{}
	for i := sl.level - 1; i >= 0; i-- {
		update[i] = sl.head
		for nil != cur.forward[i] {
			if cur.forward[i].score == score && cur.forward[i].v == v {
				update[i] = cur
				break
			}
			cur = cur.forward[i]
		}
	}

	cur = update[0].forward[0]

	for i := cur.level - 1; i >= 0; i-- {
		if update[i] == sl.head && cur.forward[i] == nil {
			sl.level = i
		}

		if nil == update[i].forward[i] {
			update[i].forward[i] = nil
		} else {
			update[i].forward[i] = update[i].forward[i].forward[i]
		}
	}

	sl.length--

	return 0
}
