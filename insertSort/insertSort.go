package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr, err := generateRandomArray(10, -50, 100)
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
			swap(arr, j, j+1)
		}
	}
}

// swap 交换数组中两个位置的值
func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
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
