package _14_sort

import "math"

//计数排序
func CountingSort(arr []int, n int) {
	if n <= 1 {
		return
	}

	//找出切片最大值
	max := math.MinInt32
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}

	//创建基数桶
	c := make([]int, max+1)
	//开始给桶计数
	for i := range arr {
		c[arr[i]]++
	}

	//桶顺序求和
	for i := 1; i <= max; i++ {
		c[i] += c[i-1]
	}

	//创建临时切片,存放数据
	r := make([]int, n)

	//从后往前遍历计数桶中数据
	for i := n - 1; i >= 0; i-- {
		index := c[arr[i]] - 1
		r[index] = arr[i]
		c[arr[i]]--
	}

	//重新赋值
	copy(arr, r)
}
