package _12_sort

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{11, 9, 20, 33, 24, 22, 30}
	QuickSort(arr)
	t.Log(arr)
}
