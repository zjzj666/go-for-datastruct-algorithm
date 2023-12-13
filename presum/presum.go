package main

import "fmt"

func main() {
	arr := []int{2, 3, 1, 5, -1, 6}
	fmt.Println(arr)
}

func preSumArray(arr []int) []int {
	n := len(arr)
	sum := make([]int, n)
	sum[0] = arr[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + arr[i]
	}
	return sum
}

func getSum(sum []int, l, r int) int {
	if l == 0 {
		return sum[r]
	} else {
		return sum[r] - sum[l-1]
	}

}
