package main

import "fmt"

// this question can have two ans: if arr can be modified : if arr cant be modified

// no modification in an array
func reverseArray(arr []int) []int {
	temp := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		temp = append(temp, arr[i])
	}
	return temp

}

//modification in an array

func reverseArrayModifies(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	arr := []int{2, 4, 6, 8}
	fmt.Println(reverseArray(arr))
	fmt.Println(reverseArrayModifies(arr))

}
