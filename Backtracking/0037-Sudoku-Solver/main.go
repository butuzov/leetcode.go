package main

import "fmt"

func main() {
	// dude, write some code...
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(board)

	for _, v := range board {
		fmt.Printf("%s\n", v)
	}
}

func solveSudoku(board [][]byte) {
	isPossible := func(row, col int, val byte) bool {
		for i := 0; i < 9; i++ {
			if board[row][i] == val {
				return false
			}
		}

		for i := 0; i < 9; i++ {
			if board[i][col] == val {
				return false
			}
		}

		//
		qx := (row / 3) * 3
		qy := (col / 3) * 3
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if board[qx+x][qy+y] == val {
					return false
				}
			}
		}

		return true
	}

	ib := map[int]byte{1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7', 8: '8', 9: '9'}

	var found bool
	var solver func()
	solver = func() {
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				if board[row][col] == '.' {
					for i := 1; i <= 9; i++ {
						if isPossible(row, col, ib[i]) {
							board[row][col] = ib[i]
							solver()

							if found {
								return
							}

							board[row][col] = '.'
						}
					}
					return
				}
			}
		}

		found = true
	}

	// run solver.
	solver()
}
