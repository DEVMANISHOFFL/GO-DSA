package main

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	currMax := nums[0]
	currMin := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// If num is negative, swap currMax and currMin
		if num < 0 {
			currMax, currMin = currMin, currMax
		}

		currMax = max(num, currMax*num)
		currMin = min(num, currMin*num)

		result = max(result, currMax)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, -2, -2, 4, -5}
	fmt.Println(maxProduct(arr))
}
