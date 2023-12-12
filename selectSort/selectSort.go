package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//随机生成一个数组并打印
	arr, err := generateRandomArray(10, -50, 100)
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
		swap(arr, i, minIndex)
	}
}

// swap 交换数组中两个位置的值
func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
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

// generateRandomArray 随机生成指定长度和范围的数组
func generateRandomArray(size, min, max int) ([]int, error) {
	//输入参数检测
	if size <= 0 || max <= min {
		return nil, errors.New("输入的长度或范围有误")
	}

	//设计随机数种子
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(max-min+1) + min // 生成[min, max]范围内的随机整数
	}
	return arr, nil
}
