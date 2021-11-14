package main

import (
	"container/heap"
	"sort"
)

func main() {
	// dude, write some code...
	println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) // expecting 4
}

// ----
// Heap Based
// -----

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := IntHeap(nums)
	heap.Init(&h)

	var s int
	for i := 0; i < k; i++ {
		s = heap.Pop(&h).(int)
	}

	return s
}

// O(n*Log(n))
func findKthLargestSort(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
