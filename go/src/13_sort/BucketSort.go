package _13_sort

import (
	_12_sort "12_sort"
	"fmt"
	"math"
)

func BucketSort(arr []int, bucketNum, maxPerBucketSize int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}
	//求出数组中最大值
	max := arr[0]
	for i := 1; i < len; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	//求出数组最小值
	min := arr[0]
	for i := 1; i < len; i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}
	//求出每个桶中得数据范围
	bucketRange := (max - min + 1) / bucketNum
	//定义桶数组
	bucketArr := make([][]int, bucketNum+1)
	index := 0
	//将数据分到对应桶中
	for i := 0; i < len; i++ {
		index = int(math.Floor(float64((arr[i] - min) / bucketRange)))
		bucketArr[index] = append(bucketArr[index], arr[i])
	}
	//将每个桶中的数据排序
	for k, v := range bucketArr {
		//桶排序
		_12_sort.QuickSort(v)
		fmt.Printf("%d --- %v\n", k, v)
	}
	return arr
}
