package main

import (
	"sort"
)

func main() {
	// dude, write some code...
	nums := [][]int{{1, 3}, {2, 4}, {10, 11}, {10, 12}, {8, 9}}
	print(minimumEffort(nums))
}

func minimumEffort(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return (tasks[i][1] - tasks[i][0]) < (tasks[j][1] - tasks[j][0])
	})

	var energy int
	for _, v := range tasks {
		energy = max(energy+v[0], v[1])
	}

	return energy
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
