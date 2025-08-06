package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func kLargest(grid [][]int, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	for i := range grid {
		for j := 0; j < len(grid[i]); j++ {
			heap.Push(h, grid[i][j])
			if h.Len() > k {
				heap.Pop(h)
			}
		}
	}
	return (*h)[0]
}

func main() {
	grid := [][]int{
		{5, 2},
		{1, 6},
	}

	fmt.Println(kLargest(grid, 3))

}
