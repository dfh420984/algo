package sort

//BucketSort 桶排序
func BucketSort(arr []int, size int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	max := getMax(arr)
	index := getBucket(max, size) //桶的数量
	bucketNum := index + 1
	bucket := make([][]int, bucketNum)
	//将数据存放桶中
	for i := 0; i < length; i++ {
		index := getBucket(arr[i], size)
		bucket[index] = append(bucket[index], arr[i])
	}
	//将桶中的数据按照从小到大排序
	tmpPos := 0
	for i := 0; i < bucketNum; i++ {
		if len(bucket[i]) >= 1 {
			QuickSort(bucket[i])
			//将桶中排序好的值放回原数组
			copy(arr[tmpPos:], bucket[i])
			tmpPos += len(bucket[i])
		}
	}
}

//getBucket
func getBucket(val int, size int) int {
	return val / size
}

func getMax(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
