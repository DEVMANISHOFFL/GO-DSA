package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {

	st, end := 0, len(arr)-1

	for st <= end {
		mid := st + (end-st)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			st = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{2, 4, 6, 8, 10, 11, 14, 16}
	fmt.Println(binarySearch(arr, 8))

}
