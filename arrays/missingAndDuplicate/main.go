package main

import "fmt"

func missingAndDuplicate(grid [][]int) []int {
	n := len(grid)
	seen := make(map[int]bool)
	actualSum := 0
	var repeated, missing int

	for i := range grid {
		for j := range grid {
			nums := grid[i][j]
			actualSum += nums
			if seen[nums] {
				repeated = nums
			}
			seen[nums] = true
		}
	}
	total := n * n
	expSum := (total * (total + 1)) / 2
	missing = expSum - (actualSum - repeated)

	return []int{repeated, missing}
}

func main() {
	arr := [][]int{{9, 2, 3}, {4, 5, 6}, {7, 7, 1}}
	fmt.Println(missingAndDuplicate(arr))

}
