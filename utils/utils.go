package utils

import (
	"errors"
	"math/rand"
	"time"
)

// Swap 交换数组中两个位置的值
func Swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

// GenerateRandomArray 随机生成指定长度和范围的数组
func GenerateRandomArray(size, min, max int) ([]int, error) {
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
