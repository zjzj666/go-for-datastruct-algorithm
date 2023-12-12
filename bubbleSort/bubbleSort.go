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

	bubbleSortMinLeft(arr)

	//输出排序后的数组
	fmt.Println(arr)
}

// bubbleSortMaxRight 冒泡排序（最大数在最右边）
func bubbleSortMaxRight(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	n := len(arr)
	for end := n - 1; end > 0; end-- {
		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
			}
		}
	}
}

// bubbleSortMinLeft 冒泡排序(最小数在最左边)
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
				swap(arr, i, i-1)
			}
		}
	}
}

// swap 交换数组中两个数的位置
func swap(arr []int, i, j int) {
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
