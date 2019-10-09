package _15_binarysearch

import "testing"

//二分查找
func TestBinarySearch(t *testing.T) {
	var arr []int = []int{10, 23, 55, 65, 78, 99}
	mid := BinarySearch(arr, 65)
	t.Log(mid)
}

func TestBinarySearchRecursive(t *testing.T) {
	var arr []int = []int{10, 23, 55, 65, 78, 99}
	mid := BinarySearchRecursive(arr, 65)
	t.Log(mid)
}
