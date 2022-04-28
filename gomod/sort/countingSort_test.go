package sort

import "testing"

func TestCountingSort(t *testing.T) {
	arr := []int{1, 5, 3, 7, 2, 4, 6}
	CountingSort(arr)
	t.Log(arr)
}
