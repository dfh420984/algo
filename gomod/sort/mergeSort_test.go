package sort

import "testing"

func TestMergeSort(t *testing.T) {
	arr := []int{1, 5, 3, 7, 2, 4, 6}
	BubbleSort(arr)
	t.Log(arr)
}
