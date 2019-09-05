package _12_sort

func QuickSort(arr []int) {
	len := len(arr)
	if len <= 1 {
		return
	}
	quickSort(arr, 0, len-1)
}

//快速排序公式 quickSort(l, q-1) quickSort(q+1, r) 终止条件:l >= r
func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	//寻找基准点，小于基准点放左边，大于放右边，然后返回基准点下标
	q := partions(arr, l, r)
	//根据基准点继续分区
	quickSort(arr, l, q-1)
	quickSort(arr, q+1, r)
}

func partions(arr []int, l, r int) int {
	//将最后一元素设为基准点
	p := arr[r]
	//定义两个游标，j从左往右依次和基准点值比较
	i := l
	j := l
	for ; j < r; j++ {
		if arr[j] < p {
			if i != j { //交换数据
				temp := arr[j]
				arr[j] = arr[i]
				arr[i] = temp
			}
			i++
		}
	}
	//比较完毕后，和初始基准点值进行数据交换
	arr[i], arr[r] = arr[r], arr[i]
	return i
}
