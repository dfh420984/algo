package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 5, 3, 7, 2, 4, 6}
	BubbleSort(arr)
	t.Log(arr)
}

func TestInsertSort(t *testing.T) {
	arr := []int{1, 5, 3, 7, 2, 4, 6}
	InsertSort(arr)
	t.Log(arr)
}

func TestSelectSort(t *testing.T) {
	arr := []int{1, 5, 3, 7, 2, 4, 6}
	SelectSort(arr)
	t.Log(arr)
}
