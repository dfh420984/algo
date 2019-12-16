package _28_heap

type Heap struct {
	a     []int
	n     int
	count int
}

//init Heap
func NewHeap(capacity int) *Heap {
	return &Heap{
		a:     make([]int, capacity+1), //用数组存储节点时，是从下标1开始，0位空着
		n:     capacity,
		count: 0,
	}
}

//大顶堆，对任意节点，子节点要比父节点值小，从下而上堆化
func (heap *Heap) Insert(data int) {
	//判断堆是否已满
	if heap.n == heap.count {
		return
	}

	//插入数据
	heap.count++
	heap.a[heap.count] = data

	//和父节点对比数据,交换
	i := heap.count
	parent := i / 2
	for parent > 0 && heap.a[parent] < heap.a[i] {
		swap(heap.a, parent, i)
		i = parent
		parent = i / 2
	}
}

//移除对顶元素，从上而下堆化
func (heap *Heap) RemoveMax() {
	if heap.count == 0 {
		return
	}

	//交换堆顶和最后一个元素
	swap(heap.a, 1, heap.count)
	//删除最后一个元素
	heap.a = heap.a[:heap.count]
	heap.count--
	//开始堆化
	heapifyUpToDown(heap.a, heap.count)
}

//从上而下堆化
func heapifyUpToDown(a []int, count int) {
	for i := 1; i <= count/2; {
		maxIndex := i
		if a[i] < a[2*i] { //如果父节点小于左子树
			maxIndex = i * 2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}
		swap(a, i, maxIndex)
		i = maxIndex
	}
}

//交换数据
func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}
