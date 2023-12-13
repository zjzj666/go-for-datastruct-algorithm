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

	bubbleSort(arr)

	//输出排序后的数组
	fmt.Println(arr)
}

// bubbleSort 冒泡排序
func bubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	n := len(arr)
	for end := n - 1; end > 0; end-- {
		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1] {
				utils.Swap(arr, i, i+1)
			}
		}
	}
}

// bubbleSortMinLeft 冒泡排序
func bubbleSortMinLeft(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	n := len(arr)
	for begin := 0; begin < n-1; begin++ {
		//比较0到n-1上的数据
		for i := n - 1; i > 0; i-- {
			//小的数据往左移动
			if arr[i] < arr[i-1] {
				utils.Swap(arr, i, i-1)
			}
		}
	}
}
