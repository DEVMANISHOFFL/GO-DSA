package main

import "fmt"

func permutation(arr []int) []int {
	piv := -1
	n := len(arr)

	//finding the pivot
	for i := n - 2; i >= 0; i-- {
		if arr[i] < arr[i+1] {
			piv = i
			break
		}
	}
	//if pivot not found -----> just reverse the array and we get next permutation here only
	if piv == -1 {
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		return arr
	}

	//2nd step---------> next larger element than pivot

	for i := n - 1; i > piv; i-- {
		if arr[i] > arr[piv] {
			arr[i], arr[piv] = arr[piv], arr[i]
			break
		}
	}

	//3rd step ---------> reverse the array from piv + 1 to n-1

	for i, j := piv+1, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

// func reverse(arr []int, start, end int) {
// 	for start < end {
// 		arr[start], arr[end] = arr[end], arr[start]
// 		start++
// 		end--
// 	}
// }

func main() {
	arr := []int{3, 2, 1}
	fmt.Println(permutation(arr))

}
