package main

import "fmt"

func search(arr []int, target int) int {
	st, end := 0, len(arr)-1

	for st <= end {

		mid := st + (end-st)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] >= arr[st] {
			if arr[st] <= target && target <= arr[mid] { // for left side -----> target start se bada aur mid se chota
				end = mid - 1
			} else {
				st = mid + 1
			}
		} else {
			if arr[mid] <= target && target <= arr[end] { // for right side-----> target mid se bada and end se chota
				st = mid + 1
			} else {
				end = mid - 1
			}
		}

	}
	return -1

}

func main() {
	arr := []int{5, 6, 7, 8, 1, 2, 3, 4, 5, 6}

	fmt.Println(search(arr, 5))
}
