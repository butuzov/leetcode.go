package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// dude, write some code...
	kth := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kth.Add(3), 4)
	fmt.Println(kth.Add(5), 5)
	fmt.Println(kth.Add(10), 5)
	fmt.Println(kth.Add(9), 8)
	fmt.Println(kth.Add(4), 8)
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	len  int
	nums *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	var kth KthLargest
	kth.len = k

	// pointing numbers int kth
	numbers := IntHeap(append([]int{}, nums...))
	kth.nums = &numbers
	heap.Init(kth.nums)

	for len(*kth.nums) > kth.len {
		heap.Pop(kth.nums)
	}

	return kth
}

func (kth *KthLargest) Add(val int) int {
	heap.Push(kth.nums, val)
	for len(*kth.nums) > kth.len {
		heap.Pop(kth.nums)
	}

	return (*kth.nums)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
