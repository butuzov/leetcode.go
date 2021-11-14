package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findRelativeRanks([]int{1, 2, 3, 4, 5}))
}

func findRelativeRanks(score []int) []string {
	answers := make([]string, 0, len(score))
	lookup := make(map[int]string, len(score))
	prized := map[int]string{
		0: "Gold Medal",
		1: "Silver Medal",
		2: "Bronze Medal",
	}

	ranked := append([]int{}, score...)

	r := ranks(ranked)
	heap.Init(&r)

	idx := 0
	for r.Len() > 0 {
		nth := heap.Pop(&r).(int)
		idx++

		if v, ok := prized[idx-1]; ok {
			lookup[nth] = v
			continue
		}
		lookup[nth] = strconv.Itoa(idx)

	}

	for i := 0; i < len(score); i++ {
		answers = append(answers, lookup[score[i]])
	}

	return answers
}

// An ranks is a min-heap of ints.
type ranks []int

func (h ranks) Len() int           { return len(h) }
func (h ranks) Less(i, j int) bool { return h[i] > h[j] }
func (h ranks) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ranks) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *ranks) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
