package main

import (
	"fmt"
	"sort"
)

func slidingWindow(arr []int, m int) int {
	sort.Ints(arr)
	n := len(arr)

	if n == 0 || n < m {
		return 0
	}

	minDiff := arr[m-1] - arr[0]

	for i := 1; i <= n-m; i++ {
		diff := arr[i+m-1] - arr[i]
		if diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}

func main() {
	arr := []int{12, 4, 7, 9, 2, 23, 25, 41, 30, 40, 28, 42, 30, 44, 48, 43, 50}
	fmt.Println(slidingWindow(arr, 7))
}
