package main

import "container/heap"

func maxSlidingWindowHeap(nums []int, k int) []int {

	if len(nums) < k || k == 0 {
		return []int{}
	}

	var maximums = make([]int, len(nums)-k+1)

	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < k-1; i++ {
		heap.Push(h, nums[i])
	}

	for i := 0; i <= len(nums)-k; i++ {
		if i > 0 {
			h.RemoveByValue(nums[i-1])
		}
		heap.Push(h, nums[k+i-1])
		maximums[i] = (*h)[0]
	}

	return maximums
}

// -- Heap Interface
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

// -- Helper
func (h *IntHeap) RemoveByValue(x int) {
	// h = *h
	for i := 0; i < len((*h)); i++ {
		if x == (*h)[i] {
			(*h) = append((*h)[:i], (*h)[i+1:]...)
			break
		}
	}

	s := *h
	heap.Init(&s)
	*h = s
}
