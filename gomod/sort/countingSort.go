package sort

//CountingSort 计数排序
func CountingSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	max := getMax(arr)
	bucket := make([]int, max+1) //桶的数量
	//桶下标存储原始元素值，遍历原数组给对应桶中计数
	for i := 0; i < length; i++ {
		bucket[arr[i]]++
	}
	//将桶中的值前后相加可以得到小于等于桶下标原始元素值的元素个数
	for i := 1; i <= max; i++ {
		bucket[i] += bucket[i-1]
	}
	r := make([]int, length)
	for i := 0; i < length; i++ {
		index := bucket[arr[i]] - 1
		r[index] = arr[i]
		bucket[arr[i]]--
	}
	copy(arr, r)
}
