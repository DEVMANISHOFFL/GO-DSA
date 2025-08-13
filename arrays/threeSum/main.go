package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums) // sort first
	n := len(nums)
	res := [][]int{}

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}

	}
	return res
}

func main() {
	nums := []int{-1, 0, 2, -2, -1, -4}
	fmt.Println(threeSum(nums)) // [[-1 -1 2] [-1 0 1]]

}
