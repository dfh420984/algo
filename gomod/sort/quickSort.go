package sort

func QuickSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	quickSort(arr, 0, length-1)
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := partition(arr, start, end)
	quickSort(arr, start, mid-1)
	quickSort(arr, mid+1, end)
}

func partition(arr []int, start, end int) int {
	pivot := arr[end] //取最后一个数作为比较点
	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(i == j) {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
