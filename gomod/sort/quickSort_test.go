package sort

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{1, 5, 3, 7, 2, 4, 6}
	QuickSort(arr)
	t.Log(arr)
}
