package grains

// package main

import (
	"errors"
	"fmt"
	"math"
)

// Total is Shorthand formula for sum (notation) of powers
// https://en.wikipedia.org/wiki/Wheat_and_chessboard_problem
func Total() uint64 {
	return (2 << 63) - 1
}

// Square using a bit opperations in order to
// do the work quickly.
func Square(n int) (uint64, error) {

	switch {
	case n < 1 || n > 64:
		return 0, errors.New("out of range")
	case n == 1:
		return 1, nil
	default:
		return 2 << uint64(n-2), nil
	}
}

// Tatal_Naive - is naive implementation of summing powers of two.
func Total_Naive() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		sq, _ := Square_Naive(i)
		sum += sq
	}
	return sum
}

// Square_Naive implements a naive approach of
// calculating grains using math.Pow
func Square_Naive(n int) (uint64, error) {

	if n < 1 || n > 64 {
		return 0, errors.New("out of range")
	}

	return uint64(math.Pow(2, float64(n-1))), nil
}

// for cmd+r tests only
// main testing (jsut for local development)
func main() {
	actualVal, actualErr := Square(1)
	fmt.Println(actualVal, actualErr)

	fmt.Printf("% 21d total good\n", Total_Naive())
	fmt.Printf("% 21d total new\n", Total())
}
