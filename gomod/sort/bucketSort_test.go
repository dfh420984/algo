package sort

import "testing"

func TestBuketSort(t *testing.T) {
	arr := []int{1, 5, 3, 7, 2, 4, 6}
	BucketSort(arr, 2)
	t.Log(arr)
}
