package main

import "fmt"

func twoSum(arr []int, target int) [][]int {
	foundingPairs := [][]int{}
	seen := make(map[int]int)

	for i, num := range arr {
		wanted := target - num
		if j, found := seen[wanted]; found {
			foundingPairs = append(foundingPairs, []int{i, j})
		}
		seen[num] = i
	}
	return foundingPairs
}

func main() {
	arr := []int{12, 7, -3, 15, 1, 8, 2, -1, 9, 5, 6, 0, 4, -2, 11}
	arr2 := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	fmt.Println(twoSum(arr, 10))
	fmt.Println(twoSum(arr2, 10))
}
