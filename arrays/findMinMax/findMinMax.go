package main

import "fmt"

func findMinMax(arr []int) (int, int) {
	n := len(arr)
	var min, max, i int

	if n == 0 {
		return 0, 0
	}

	if n%2 == 0 {
		if arr[0] < arr[1] {
			min = arr[0]
			max = arr[1]
		} else {
			max = arr[0]
			min = arr[1]
		}
		i = 2
	} else {
		min = arr[0]
		max = arr[0]
		i = 1
	}

	for i < n-1 {
		if arr[i] < arr[i+1] { //arr[i+1] bada hai
			if arr[i+1] > max { // arr[i+1] bada h to kya wo max se bhi bada hai
				max = arr[i+1] // if hai to max is now updated
			}
			if arr[i] < min { //similarly arr[i] chota hai to kya wo min se bhi chota hai
				min = arr[i] //if hai to update krdo
			}
		} else { //opposite case-------------> arr[i] bada hai
			if arr[i] < max {
				max = arr[i]
			}
			if arr[i+1] < min {
				min = arr[i+1]
			}

		}
		i += 2

	}
	return min, max
}
func main() {
	arr := []int{45, 2, 999, -23, 76, 5, 0, -999, 435, 12, 67, -34, 888, 3, 1001}
	min, max := findMinMax(arr)
	fmt.Println("min:", min)
	fmt.Println("max:", max)

}
