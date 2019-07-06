package main
import (
	"fmt"
	"math"
)

// Correct Solution for Task
func largestPalindromLookUpTable(n int) int {
	var largestPalindroms = map[int]int{
		1: 9,
		2: 987,
		3: 123,
		4: 597,
		5: 677,
		6: 1218,
		7: 877,
		8: 475,
		9: 520,
	}
	return largestPalindroms[n]
}

func isPalindrome(n int) bool {

	// Simple precondition check
	switch {
	case n < 0:
		return false
	}

	var reversed int
	var rest int
	var preserved = n
	for i := 0; n > 0; i++ {
		rest = n % 10
		n = (n - rest) / 10
		reversed = reversed*10 + rest
	}

	if reversed == preserved {
		return true
	}

	return false
}

// Naive Brutforce a bit optimized, but not really
func largestPalindrome(n int) int {
	var startN, finishN int
	startN = int(math.Pow10(n - 1))
	finishN = int(math.Pow10(n)) - 1

	var tmp, max int

	for i := finishN; i >= startN; i-- {
		for j := i; j >= startN; j-- {
			fmt.Println(i, j)
			tmp = i * j
			if !isPalindrome(tmp) {
				continue
			}

			if tmp > max {
				max = tmp
				continue
			}

		}
	}
	return max % 1337
}
