package main

import "fmt"

func main() {
	// dude, write some code...
	b := [][]byte{
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

	fmt.Println(isValidSudoku(b))
}

func isValidSudoku(board [][]byte) bool {
	isPossible := func(row, col int, val byte) bool {
		var rows int
		for i := 0; i < 9; i++ {
			if board[row][i] == val {
				rows++
			}
		}
		if rows > 1 {
			return false
		}

		var cols int
		for i := 0; i < 9; i++ {
			if board[i][col] == val {
				cols++
			}
		}
		if cols > 1 {
			return false
		}

		//
		var quandrant int

		qx := (row / 3) * 3
		qy := (col / 3) * 3
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if board[qx+x][qy+y] == val {
					quandrant++
				}
			}
		}

		if quandrant > 1 {
			return false
		}

		return true
	}

	//

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}

			if !isPossible(row, col, board[row][col]) {
				return false
			}

		}
	}

	return true
}
