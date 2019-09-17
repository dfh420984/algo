package sort

//冒泡排序
func BubbleSort(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}
	for i := 0; i < len; i++ {
		flag := false
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}

//插入排序
func InsertSort(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}
	for i := 1; i < len; i++ {
		val := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > val {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = val
	}
	return arr
}

//选择排序
func SelectSort(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}
	for i := 0; i < len; i++ {
		min := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			tmp := arr[min]
			arr[min] = arr[i]
			arr[i] = tmp
		}
	}
	return arr
}
