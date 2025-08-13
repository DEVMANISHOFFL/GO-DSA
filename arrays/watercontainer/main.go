package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func container(arr []int) int {
	l, r := 0, len(arr)-1
	maxArea := 0

	for l < r {
		// Calculate area
		height := 0
		height = min(arr[l], arr[r])

		area := height * (r - l)
		maxArea = max(maxArea, area)

		// Move the smaller height pointer inward
		if arr[l] < arr[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(container(arr)) // ✅ Output: 49
}
