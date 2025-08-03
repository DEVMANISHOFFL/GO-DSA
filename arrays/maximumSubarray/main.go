package main

import "fmt"

func maxSubArraySum(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	maxSum := arr[0]

	for st := range arr {
		currSum := 0
		for end := st; end < n; end++ {
			currSum = currSum + arr[end]
			if currSum > maxSum {
				maxSum = currSum
			}
		}
	}
	return maxSum
}

func main() {
	arr := []int{12, -5, 7, -3, 15, -20, 8, 9, -1, 0, 6, -4, 10, -2, 11, -30, 25}
	// arr := []int{1, 3, 5, 7}
	fmt.Println(maxSubArraySum(arr))
}
