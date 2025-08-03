package main

import "fmt"

func twoSum(arr []int, target int) [][]int {
	foundingPairs := [][]int{}
	seen := make(map[int]int)

	for i, num := range arr {

		toFind := target - num
		if j, found := seen[toFind]; found {
			fmt.Println(j)
			foundingPairs = append(foundingPairs, []int{i, j})
		}
		seen[num] = i
		// fmt.Println(seen)
	}
	return foundingPairs
}

func main() {
	arr := []int{1, 3, 5, 7, 2, 4, 6, 8}
	fmt.Println(twoSum(arr, 10))
}
