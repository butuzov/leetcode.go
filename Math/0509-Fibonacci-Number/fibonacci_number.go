package main
import (
	"math"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/

// fib is naive implementation of Fibonachi Sequence
func fibonacci(N int) int {
	if N < 2 {
		return N
	}

	return fibonacci(N-2) + fibonacci(N-1)
}

// Binet's Formula for quick calculation of Fibonachi
func fibonacciBinet(N int) int {
	p1 := 1 / math.Pow(5, .5)
	p2 := math.Pow((1+math.Pow(5, .5))/2, float64(N))
	p3 := math.Pow((1-math.Pow(5, .5))/2, float64(N))
	return int(math.Round(p1 * (p2 - p3)))
}

// Swaping Technics

func fibonacciSwaps(N int) int {
	a, b := 0, 1

	for i := 0; i < N; i++ {
		a, b = b, a+b
	}

	return a
}

// Memoization technics with naive implementation

var fibonacciMemoize func(int) int

func memoize(fnc func(int) int) func(int) int {
	var cache = make(map[int]int)

	return func(N int) int {
		if val, ok := cache[N]; ok {
			return val
		}
		cache[N] = fnc(N)
		return cache[N]
	}
}

func init() {
	fibonacciMemoize = memoize(func(N int) int {
		if N < 2 {
			return N
		}
		return fibonacciMemoize(N-2) + fibonacciMemoize(N-1)
	})

}
