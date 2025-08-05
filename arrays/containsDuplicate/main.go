package main

import (
	"fmt"
	"sort"
)

// func bruteforceDuplication(nums []int) bool {
// 	for i := range nums {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] == nums[j] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

func withNoExtraSpaceCheck(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func containsDuplicates(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

func main() {
	arr := []int{1, 3, 5, 7, 345, 45, 34, 356, 365}
	arr2 := []int{1, 3, 5, 7, 345, 45, 34, 356, 3, 3, 365}
	// fmt.Println(bruteforceDuplication(arr))
	fmt.Println(containsDuplicates(arr))
	fmt.Println(withNoExtraSpaceCheck(arr2))

}
