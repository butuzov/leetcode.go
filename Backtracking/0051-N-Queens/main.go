package main

import (
	"fmt"
)

func main() {

	for _, b := range solveNQueens(4) {
		fmt.Println(b)
		fmt.Println()
	}

}
