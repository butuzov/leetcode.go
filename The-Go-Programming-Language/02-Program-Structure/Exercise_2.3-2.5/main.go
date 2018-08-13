package main

import (
	"fmt"

	pc "./popcount"
)

func main() {
	var x uint64 = 7
	fmt.Printf("%.8b - %d\n", x, x)
	fmt.Printf("%.8b (%d) - %d\n", x, x, pc.PopCount(x))
	fmt.Printf("%.8b (%d) - %d\n", x, x, PopCountByClearing(x))
	fmt.Printf("%.8b (%d) - %d\n", x, x, PopCountByShifting(x))
}

func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}
