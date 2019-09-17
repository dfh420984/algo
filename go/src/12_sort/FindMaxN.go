package _12_sort

//寻找无序数组中第K大元素
func FindMaxN(arr []int, n int) int {
	len := len(arr)
	if len < n || n <= 0 {
		return 0
	}
	//寻找分区点
	p := kpartions(arr, 0, len-1)
	for p+1 != n {
		if p+1 > n { //在分区左边边查找
			p = kpartions(arr, 0, p-1)
		} else {
			p = kpartions(arr, p+1, len-1)
		}
	}
	return arr[p]
}

//分区数据按从大到小排序
func kpartions(arr []int, l, r int) int {
	if l >= r {
		return r
	}
	//用最后一个元素作为比较基准点
	p := arr[r]
	i := l
	j := l
	for ; j < r; j++ {
		if arr[j] > p {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[r], arr[i] = arr[i], arr[r]
	return i
}
