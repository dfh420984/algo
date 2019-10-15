package main

import (
	"container/list"
	"fmt"
)

//adjacency table (邻接表), 无向图
type Graph struct {
	adj []*list.List //双向链表
	v   int          //节点个数
}

func NewGraph(v int) *Graph {
	graph := &Graph{}
	graph.v = v
	graph.adj = make([]*list.List, v)
	for i := range graph.adj {
		graph.adj[i] = list.New()
	}
	return graph
}

//insert as add edge，一条边存2次
func (this *Graph) addEdge(s int, t int) {
	this.adj[s].PushBack(t)
	this.adj[t].PushBack(s)
}

//search path by BFS (广度优先搜索)
func (this *Graph) bfs(s int, t int) {
	if s == t {
		return
	}
	//访问路径
	pre := make([]int, this.v)
	for i := range pre {
		pre[i] = -1
	}
	var queue []int
	visited := make([]bool, this.v)
	queue = append(queue, s)
	visited[s] = true
	isFound := false
	for len(queue) > 0 && !isFound {
		top := queue[0]
		queue = queue[1:]
		linklist := this.adj[top]
		for e := linklist.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			if !visited[k] {
				pre[k] = top
				if k == t {
					isFound = true
					break
				}
				queue = append(queue, k)
				visited[k] = true
			}
		}
	}
	if isFound {
		this.printPre(pre, s, t)
	} else {
		fmt.Printf("no path found from %d to %d\n", s, t)
	}
}

//深度优先搜索
func (this *Graph) dfs(s int, t int) {
	pre := make([]int, this.v)
	for i := range pre {
		pre[i] = -1
	}
	visited := make([]bool, this.v)
	visited[s] = true
	isFound := false
	this.recurse(s, t, visited, pre, isFound)
	this.printPre(pre, s, t)
}

//递归
func (this *Graph) recurse(s int, t int, visited []bool, pre []int, isFound bool) {
	if isFound {
		return
	}
	visited[s] = true
	if s == t {
		isFound = true
		return
	}
	linklist := this.adj[s]
	for e := linklist.Front(); e != nil; e = e.Next() {
		k := e.Value.(int)
		if !visited[k] {
			pre[k] = s
			this.recurse(k, t, visited, pre, isFound)
		}
	}
}

func (this *Graph) printPre(pre []int, s int, t int) {
	if t != s && pre[t] != -1 {
		this.printPre(pre, s, pre[t])
	}
	fmt.Printf("%d ", t)

}

func main() {
	graph := NewGraph(8)
	graph.addEdge(0, 1)
	graph.addEdge(0, 3)
	graph.addEdge(1, 2)
	graph.addEdge(1, 4)
	graph.addEdge(2, 5)
	graph.addEdge(3, 4)
	graph.addEdge(4, 5)
	graph.addEdge(4, 6)
	graph.addEdge(5, 7)
	graph.addEdge(6, 7)

	//graph.bfs(0, 6)
	graph.dfs(0, 7)
	fmt.Println()
	graph.dfs(1, 3)
}
