package _13_sort

import "testing"

func TestBucketSort(t *testing.T) {
	arr := []int{11, 23, 45, 67, 88, 99, 22, 34, 56, 78, 90, 12, 34, 5, 6}
	BucketSort(arr, 4, 4)
}
