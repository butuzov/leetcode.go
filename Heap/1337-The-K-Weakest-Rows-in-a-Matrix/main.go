package main

import (
	"container/heap"
	"fmt"
)

type row struct {
	idx int
	rnk int
	row []int
}

type Rows []row

func (h *Rows) Len() int { return len(*h) }
func (h *Rows) Less(i, j int) bool {
	if (*h)[i].rnk == (*h)[j].rnk {
		return (*h)[i].idx > (*h)[j].idx
	}

	return (*h)[i].rnk > (*h)[j].rnk
}
func (h *Rows) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *Rows) Push(x interface{}) { *h = append(*h, x.(row)) }

func (h *Rows) Pop() interface{} {
	old, n := h, len(*h)
	x := (*old)[n-1] // last element
	*h = (*old)[0 : n-1]
	return x
}

func calcRank(r []int) int {
	var i int

	for _, v := range r {
		if v == 0 {
			break
		}
		i += v
	}

	return i
}

func main() {
	// dude, write some code...

	m := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}

	fmt.Println(kWeakestRows(m, 3))
}

func kWeakestRows(mat [][]int, k int) []int {
	rows := make(Rows, 0, k)
	heap.Init(&rows)

	for idx, r := range mat {
		soldiers := row{
			idx: idx,
			rnk: calcRank(r),
			row: r,
		}

		heap.Push(&rows, soldiers)
		if rows.Len() > k {
			heap.Pop(&rows)
		}

	}

	idxs := make([]int, k)
	cont := k - 1
	for rows.Len() > 0 {
		r := heap.Pop(&rows).(row)
		idxs[cont] = r.idx
		cont--
	}

	return idxs
}
