package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rain(hi []int) int {
	n := len(hi)
	l, r, lmax, rmax, ans := 0, n-1, 0, 0, 0

	for l < r {
		lmax = max(lmax, hi[l])
		rmax = max(rmax, hi[r])

		if rmax < lmax {
			ans += rmax - hi[r]
			r--
		} else {
			ans += lmax - hi[l]
			l++
		}
	}
	return ans
}

func main() {
	height := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(rain(height))
}
