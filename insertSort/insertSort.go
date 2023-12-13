package main

import (
	"fmt"
	"go-for-datastruct-algorithm/utils"
)

func main() {
	arr, err := utils.GenerateRandomArray(10, -50, 100)
	if err != nil {
		panic(err)
	}

	//输出随机生成的数组
	fmt.Println(arr)

	insertSort(arr)

	//输出排序后的数组
	fmt.Println(arr)
}

// insertSort 插入排序
func insertSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			utils.Swap(arr, j, j+1)
		}
	}
}
