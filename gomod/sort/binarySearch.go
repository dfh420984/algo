package sort

//BinarySearch 二分查找
func BinarySearch(arr []int, val int) int {
	length := len(arr)
	if length <= 0 {
		return -1
	}
	low := 0
	high := length - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == val {
			return mid
		} else if arr[mid] < val { //在中间值的右边查找
			low = mid + 1
		} else { //在中间值的左边查找
			high = mid - 1
		}
	}
	return -1
}

//BinarySearchRecursion 递归查找
func BinarySearchRecursion(arr []int, val int) int {
	return bsRecursion(arr, 0, len(arr)-1, val)
}

func bsRecursion(arr []int, low, high, val int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if arr[mid] == val {
		return mid
	} else if arr[mid] < val {
		return bsRecursion(arr, mid+1, high, val)
	} else {
		return bsRecursion(arr, low, mid-1, val)
	}
}

//变体二分查找，查找第一个等于给定值的下标
func BinarySearchFirst(arr []int, val int) int {
	length := len(arr)
	if length <= 0 {
		return -1
	}
	low := 0
	high := length - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] < val {
			low = mid + 1
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			if mid == 0 || arr[mid-1] != val {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

//变体二分查找，查找最后一个等于给定值的下标
func BinarySearchLast(arr []int, val int) int {
	length := len(arr)
	if length <= 0 {
		return -1
	}
	low := 0
	high := length - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] < val {
			low = mid + 1
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			if mid == length-1 || arr[mid+1] != val {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

//查找第一个大于给定值的元素
func BinarySearchFirstGT(arr []int, val int) int {
	length := len(arr)
	if length <= 0 {
		return -1
	}
	low := 0
	high := length - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] < val {
			low = mid + 1
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			if mid != length-1 && arr[mid+1] > val {
				return mid + 1
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

//查找最后一个小于给定值的元素
func BinarySearchFirstLT(arr []int, val int) int {
	length := len(arr)
	if length <= 0 {
		return -1
	}
	low := 0
	high := length - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] < val {
			low = mid + 1
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			if mid != 0 && arr[mid-1] < val {
				return mid - 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
