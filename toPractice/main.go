package main

import "fmt"

func maxSubArraySum(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	currSum, maxSum := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		currSum = max(arr[i], currSum+arr[i])
		maxSum = max(maxSum, currSum)
	}
	return maxSum

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	arr := []int{2, 4, 6, 8, 0, 1, 2, 3, -4, 5, 6, -7, -8, 9}
	fmt.Println(maxSubArraySum((arr)))
}
