package _12_sort

import "testing"

func TestMergerSort(t *testing.T) {
	arr := []int{11, 9, 20, 33, 24, 22, 30}
	MergerSort(arr)
	t.Log(arr)
}
