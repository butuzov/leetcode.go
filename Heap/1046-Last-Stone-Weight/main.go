package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// dude, write some code...
	fmt.Println("Result is ", lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}

// An ints is a min-heap of ints.
type ints []int

func (h *ints) Len() int           { return len(*h) }
func (h *ints) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *ints) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *ints) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *ints) Pop() interface{} {
	old, n := h, len(*h)
	x := (*old)[n-1] // last element
	*h = (*old)[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	stonesh := ints(stones)
	heap.Init(&stonesh)

	for {

		if stonesh.Len() == 1 {
			return heap.Pop(&stonesh).(int)
		}

		if stonesh.Len() == 0 {
			break
		}

		x := heap.Pop(&stonesh).(int)
		y := heap.Pop(&stonesh).(int)
		d := x - y

		if d == 0 {
			continue
		}

		if d < 0 {
			d *= -1
		}

		heap.Push(&stonesh, d)

	}

	return 0
}
