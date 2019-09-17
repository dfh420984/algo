package _12_sort

import "testing"

func TestFindMaxN(t *testing.T) {
	arr := []int{6, 3, 1, 5, 7, 4, 2, 9, 11, 8}
	res := FindMaxN(arr, 4)
	t.Log(res)
	t.Log(arr)
}
