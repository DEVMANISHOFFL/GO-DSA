package main

import "fmt"

func next(arr []int) []int {
	n := len(arr)
	//assume pivot point to be -1
	pivot := -1

	// finding pivot point ----------> ex:- [1,2,3,6,5,4] ------. pivot is 3
	for i := n - 2; i >= 0; i-- {
		if arr[i] < arr[i+1] {
			pivot = i
			break
		}
	}

	//if no pivot found just reverse the arr , thats next permutation
	if pivot == -1 {
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		return arr
	}

	//if pivot found , find number next bigger than it and swap it with the pivot point

	for i := n - 1; i > pivot; i-- {
		if arr[pivot] < arr[i] {
			arr[pivot], arr[i] = arr[i], arr[pivot]
			break
		}
	}

	//now simply reverse the string from pivot+1 to n-1--------> final answers
	for i, j := pivot+1, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

//can also make a reverse function for convnience and clean code

// func reverse(arr []int, start, end int) {
// 	for start < end {
// 		arr[start], arr[end] = arr[end], arr[start]
// 		start++
// 		end--
// 	}
// }

func main() {
	arr := []int{1, 1, 5}

	fmt.Println(next(arr))

}
