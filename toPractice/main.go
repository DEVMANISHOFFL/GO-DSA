package main

import "fmt"

func reverseArray(arr []int) []int {
	temp := []int{}

	for i := len(arr) - 1; i >= 0; i-- {
		temp = append(temp, arr[i])
	}
	return temp
}

func reverseArrayModifies(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(reverseArray(arr))
	fmt.Println(reverseArrayModifies(arr2))
}
