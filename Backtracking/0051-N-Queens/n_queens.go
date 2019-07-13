package main

import "fmt"

/*******************************************************************************
  Problem Solution
*******************************************************************************/

// Solve N Queens problem solved in a next way
// first -> generating boards with backpropogations
// second -> receiving results via channel
func solveNQueens(n int) [][]string {

	var results = [][]string{}
	if n == 0 {
		return results
	}

	var items = make(chan []int)
	var stop = make(chan bool)

	// part 1
	go func(size int, results chan<- []int, stop chan<- bool) {

		// backpropogation.
		var generate func(*[]int, int)
		generate = func(b *[]int, row int) {
			if row > size {
				var pkg = make([]int, size)
				copy(pkg, *b)
				results <- pkg
				// fmt.Println("OK", *b)
				return
			}
			if len(*b) < row {
				*b = append(*b, 0)
			}
			for i := 1; i <= size; i++ {
				if attacked(b, &i) {
					continue
				}
				(*b)[row-1] = i
				generate(b, row+1)
			}
			*b = (*b)[:len(*b)-1]

			if row == 1 {
				stop <- true
			}
		}

		var board = []int{}
		generate(&board, 1)

	}(n, items, stop)

	// part 2

	Results(items, stop, &results)

	return results
}

// Results wait for results etc...
func Results(items <-chan []int, stop <-chan bool, res *[][]string) {
	for {
		select {
		case board := <-items:
			*res = append(*res, Strings(board))
		case <-stop:
			return
		}
	}
	fmt.Println("STOP")
	return
}

// Abs value of int as helper function.
func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// attacked will check can queen at
//  row board[:-1] and colAtc to be attacked by any of
// other queens at board
func attacked(board *[]int, colAtc *int) bool {
	var rowAtc = len(*board)

	for row, col := range (*board)[:len(*board)-1] {
		if (col == *colAtc) || (rowAtc-(row+1)) == Abs(*colAtc-col) {
			return true
		}
	}

	return false
}

// Strings will convert slice of ints (for example)
// []int{1,2,3}
// into slice of strings
// []string{
// "Q.." ,
// ".Q." ,
// "..Q" ,
// }
func Strings(n []int) []string {
	var (
		board = make([]string, len(n))
		row   = make([]byte, len(n))
	)

	for i := 0; i < len(n); i++ {
		for ri := 0; ri < len(n); ri++ {

			if ri == n[i]-1 {
				row[ri] = 'Q'
				continue
			}
			row[ri] = '.'
		}

		board[i] = string(row)
	}

	return board
}
