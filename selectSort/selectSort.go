package main

import (
	"errors"
	"fmt"
	"go-for-datastruct-algorithm/utils"
)

func main() {
	//随机生成一个数组并打印
	arr, err := utils.GenerateRandomArray(10, -50, 100)
	if err != nil {
		panic(err)
	}

	//输出随机生成的数组

	fmt.Println(arr)

	selectionSort(arr)

	//输出排序后的数组
	fmt.Println(arr)
}

// selectionSort 选择排序
func selectionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex, err := findMinIndexInArray(arr, i, len(arr)-1)
		if err != nil {
			panic(err)
		}
		//把minIndex位置的数与i交换
		utils.Swap(arr, i, minIndex)
	}
}

// findMinIndexInArray 查询给定范围下数组中最小值索引
func findMinIndexInArray(arr []int, begin int, end int) (int, error) {
	if begin > end {
		return 0, errors.New("输入范围错误")
	}
	if len(arr) < 2 || begin == end {
		return begin, nil
	}
	minIndex := begin
	for i := begin + 1; i <= end; i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}
	return minIndex, nil
}
