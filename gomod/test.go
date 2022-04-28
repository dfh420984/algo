package main

import "fmt"

func main() {
	bucket := make([]int, 10) //桶的数量
	arr := []int{1, 5, 3, 7, 2, 4, 20, 6, 8, 9}
	bucket[arr[1]] = bucket[arr[1]] + 1
	fmt.Println(bucket)
}
