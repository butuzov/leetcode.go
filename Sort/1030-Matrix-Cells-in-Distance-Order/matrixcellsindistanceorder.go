package main
import (
	"container/heap"
	"fmt"
)

// heap container

type Cell struct {
	distance int // manhatten distance
	row, col int // row and cell to return
}

type CellHeap []Cell

func (c Cell) String() string { return fmt.Sprintf("(%d, %d) - %d", c.row, c.col, c.distance) }

// heap interface
func (c CellHeap) Len() int            { return len(c) }
func (c CellHeap) Less(i, j int) bool  { return c[i].distance < c[j].distance }
func (c CellHeap) Swap(i, j int)       { c[i], c[j] = c[j], c[i] }
func (c *CellHeap) Push(x interface{}) { *c = append(*c, x.(Cell)) }
func (c *CellHeap) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[0 : n-1]
	return x
}

// abs
func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {

	var results = make([][]int, R*C)
	var heapC = &CellHeap{}
	heap.Init(heapC)

	for row := 0; row < R; row++ {
		for col := 0; col < C; col++ {
			heap.Push(heapC, Cell{abs(row-r0) + abs(col-c0), row, col})
		}
	}

	var i = 0
	var cell Cell
	for heapC.Len() > 0 {
		cell = (heap.Pop(heapC)).(Cell)
		results[i] = []int{cell.row, cell.col}
	}

	return results
}
