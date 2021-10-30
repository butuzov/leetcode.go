package main

import (
	"container/heap"
	"sort"
)

func main() {
	// dude, write some code...
	println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 4))
}

type lh []int

func (h *lh) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *lh) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *lh) Len() int           { return len(*h) }
func (h *lh) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
func (h *lh) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *lh) Peak() int          { return (*h)[0] }

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	pool := make([][2]int, n)
	for i := 0; i < n; i++ {
		pool[i] = [2]int{efficiency[i], speed[i]}
	}

	sort.Slice(pool, func(i, j int) bool {
		return pool[i][0] > pool[j][0]
	})

	maxPerf, sum, h := 0, 0, new(lh)
	for _, worker := range pool {
		if h.Len() >= k {
			sum -= heap.Pop(h).(int)
		}
		sum += worker[1]
		heap.Push(h, worker[1])

		if curPerf := sum * worker[0]; curPerf > maxPerf {
			maxPerf = curPerf
		}
	}

	return maxPerf % 1_000_000_007
}
