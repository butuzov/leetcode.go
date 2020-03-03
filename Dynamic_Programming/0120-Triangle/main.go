package main

import "fmt"

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func minimumTotal(triangle [][]int) int {

	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	type index struct {
		row int
		idx int
	}
	var lookup = make(map[index]int)

	var down func(row int, idx int) int

	var memo = func(row int, idx int) int {
		var ri = index{row, idx}
		if val, ok := lookup[ri]; !ok {
			lookup[ri] = down(row, idx)
			return lookup[ri]
		} else {
			return val
		}
	}
	down = func(row int, idx int) int {

		if row == len(triangle)-1 {
			return triangle[row][idx]
		}

		return triangle[row][idx] + min(memo(row+1, idx), memo(row+1, idx+1))
	}

	return memo(0, 0)
}

func main() {

	fmt.Println(minimumTotal(TestCases[2].triangle))
}
