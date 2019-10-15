package _15_binarysearch

//二分查找循环
func BinarySearch(arr []int, val int) int {
	len := len(arr)
	if len == 0 {
		return -1
	}
	l := 0
	h := len - 1
	for l <= h {
		mid := l + (h-l)/2
		if arr[mid] > val { //查找值小于中间值，在中间值左边查找
			h = mid - 1
		} else if arr[mid] < val { //查找值大于中间值，在中间值右边查找
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

//二分查找递归
func BinarySearchRecursive(arr []int, val int) int {
	len := len(arr)
	if len == 0 {
		return -1
	}
	return bs(arr, 0, len-1, val)
}

func bs(arr []int, l, h, val int) int {
	if l > h {
		return -1
	}
	mid := l + (h-l)/2
	if arr[mid] == val {
		return mid
	} else if arr[mid] > val {
		return bs(arr, l, mid-1, val)
	} else {
		return bs(arr, mid+1, h, val)
	}
}

//查找第一个给定值的元素
func BinarySearchFirst(arr []int, val int) int {
	len := len(arr)
	if len == 0 {
		return -1
	}
	l := 0
	h := len - 1
	for l <= h {
		mid := l + (h-l)/2
		if arr[mid] > val { //查找值小于中间值，在中间值左边查找
			h = mid - 1
		} else if arr[mid] < val { //查找值大于中间值，在中间值右边查找
			l = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != val {
				return mid
			} else {
				h = mid - 1
			}
		}
	}
	return -1
}

//查找最后一个给定值的元素
func BinarySearchLast(arr []int, val int) int {
	len := len(arr)
	if len == 0 {
		return -1
	}
	l := 0
	h := len - 1
	for l <= h {
		mid := l + (h-l)/2
		if arr[mid] > val { //查找值小于中间值，在中间值左边查找
			h = mid - 1
		} else if arr[mid] < val { //查找值大于中间值，在中间值右边查找
			l = mid + 1
		} else {
			if mid == h || arr[mid+1] != val {
				return mid
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}

//查找第一个大于等于给定值的元素
func BinarySearchGT(arr []int, val int) int {
	len := len(arr)
	if len == 0 {
		return -1
	}
	l := 0
	h := len - 1
	for l <= h {
		mid := l + (h-l)/2
		if arr[mid] > val { //查找值小于中间值，在中间值左边查找
			h = mid - 1
		} else if arr[mid] < val { //查找值大于中间值，在中间值右边查找
			l = mid + 1
		} else {
			if mid != len-1 && a[mid+1] > val {
				return mid + 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}

//查找最后一个小于等于给定值的元素
func BinarySearchLT(arr []int, val int) int {
	len := len(arr)
	if len == 0 {
		return -1
	}
	l := 0
	h := len - 1
	for l <= h {
		mid := l + (h-l)/2
		if arr[mid] > val { //查找值小于中间值，在中间值左边查找
			h = mid - 1
		} else if arr[mid] < val { //查找值大于中间值，在中间值右边查找
			l = mid + 1
		} else {
			if mid == 0 && a[mid-1] < val {
				return mid - 1
			} else {
				h = mid - 1
			}
		}
	}
	return -1
}
