package sort

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11}
	t.Log(BinarySearch(arr, 3))
}

func TestBinarySearchRecursion(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11}
	t.Log(BinarySearchRecursion(arr, 9))
}
