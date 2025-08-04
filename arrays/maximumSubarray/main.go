package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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
			maxSum = max(currSum, maxSum)
		}
	}
	return maxSum
}

func kadanesAlgo(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	currSum, maxSum := arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		currSum = max(arr[i], currSum+arr[i])
		maxSum = max(currSum, maxSum)

	}
	return maxSum

}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(kadanesAlgo(arr))
	fmt.Println(maxSubArraySum(arr))
}
