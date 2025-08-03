package main

import "fmt"

func findMinMax(arr []int) (int, int) {
	var min, max, i int
	n := len(arr)

	if n == 0 {
		return 0, 0
	}

	if n%2 == 0 {
		if arr[0] > arr[1] {
			min = arr[1]
			max = arr[0]
		} else {
			max = arr[1]
			min = arr[0]
		}
		i = 2
	} else {
		min = arr[0]
		max = arr[0]
		i = 1
	}

	for i < n-1 {
		if arr[i] > arr[i+1] {
			if arr[i] > max {
				max = arr[i]
			}
			if arr[i+1] < min {
				min = arr[i+1]
			}
		} else {
			if arr[i] < min {
				min = arr[i]
			}
			if arr[i+1] > max {
				max = arr[i+1]
			}

		}
		i = i + 2
	}

	return min, max
}

func main() {
	arr := []int{45, 2, 999, -23, 76, 5, 0, -999, 435, 12, 67, -34, 888, 3, 1001}
	min, max := findMinMax(arr)
	fmt.Println("min:", min)
	fmt.Println("max:", max)

}
