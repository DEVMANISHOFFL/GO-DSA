package main

import "fmt"

func min(arr []int) int {
	high, low := len(arr)-1, 0
	for high > low {
		mid := low + (high-low)/2
		if arr[mid] > arr[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return arr[low]
}

func main() {
	arr := []int{8, 9, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(min(arr))
}
