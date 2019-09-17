package _12_sort

func MergerSort(arr []int) {
	arrlen := len(arr)
	if arrlen <= 1 {
		return
	}

	mergerSort(arr, 0, arrlen-1)
}

//归并排序 递推公式 merge(mergerSort[l...q],mergerSort[q+1...r]) l >= r
func mergerSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	//寻找中间点
	q := l + (r-l)/2

	//开始递归分解，直到每个组中只剩最后一个元素
	mergerSort(arr, l, q)
	mergerSort(arr, q+1, r)

	//开始合并，排序分解后的子数组
	merge(arr, l, q, r)
}

func merge(arr []int, l, m, r int) {
	temp := make([]int, r-l+1)
	//定义两个游标，分别指向子数组起始位置
	i := l
	j := m + 1
	k := 0

	//开始比较自分组中数据大小数据小的在前
	for ; i <= m && j <= r; k++ {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
	}

	//将子数组中剩余的元素插入临时数组中temp中
	for ; i <= m; i++ {
		temp[k] = arr[i]
		k++
	}

	for ; j <= r; j++ {
		temp[k] = arr[j]
		k++
	}

	//将临时数据中的元素替换原数组值
	for z := 0; z < len(temp); z++ {
		arr[l+z] = temp[z]
	}
}
