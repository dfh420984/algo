package sort

//BubbleSort 冒泡排序 go切片是引用类型 时间复杂度O(n^2)，原地排序，稳定排序
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		//判断本轮是否需要交换
		flag := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag { //本轮没有交换，则退出
			break
		}
	}
}

//InsertSort 插入排序 时间复杂度O(n^2)，原地排序，稳定排序
func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ { //待排序的数组
		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- { //已排序的数组
			if value < arr[j] {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = value
	}
}

//SelectSort 选择排序 时间复杂度O(n^2)，原地排序，不稳定排序
func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		//找到最小值
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		//交换本轮最小值和本轮第一个值
		arr[i], arr[min] = arr[min], arr[i]
	}
}
