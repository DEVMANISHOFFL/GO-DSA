package main

import (
	"fmt"
	"sort"
)

func slidingWindow(arr []int, m int) (int, []int) {
	sort.Ints(arr)
	n := len(arr)

	if n == 0 || n < m {
		return 0, []int{}
	}

	minDiff := arr[m-1] - arr[0]
	st, end := 0, m-1

	for i := 0; i < n-m+1; i++ {
		diff := arr[i+m-1] - arr[i]
		if diff < minDiff {
			minDiff = diff
			st, end = i, i+m-1
		}
	}
	return minDiff, arr[st : end+1]
}

func main() {
	arr := []int{12, 4, 7, 9, 2, 23, 25, 41, 30, 40, 28, 42, 30, 44, 48, 43, 50}
	fmt.Println(slidingWindow(arr, 8))
}
