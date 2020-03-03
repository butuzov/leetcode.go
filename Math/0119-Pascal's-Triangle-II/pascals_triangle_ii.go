package main

import "fmt"

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func getRow(n int) []int {
	var _row = make([]int, n+1)

	for k := 0; k < (n/2)+1; k++ {
		_row[k] = Idx(n, k)
		_row[n-k] = _row[k]
	}

	return _row
}

// Idx used binominal coefficient, but in order to avoid overflow
// we are cutting out factorials for n and max(k, n-k). And its not
// stable with int, overflow is comming.
func Idx(n, k int) int {

	var max, min int
	if k > n-k {
		max = k
		min = n - k
	} else {
		max = n - k
		min = k
	}

	var (
		divisible = 1
		diviser   = fac(min)
	)
	for i := n; i > max; i-- {
		if diviser%i == 0 {
			diviser /= i
		} else {
			divisible *= i
		}
	}

	return divisible / diviser
}

// Factorial calculator with a simple memoizer
var factorials = map[int]int{0: 1, 1: 1}

func fac(n int) int {
	if _, ok := factorials[n]; !ok {
		factorials[n] = n * fac(n-1)
	}
	return factorials[n]
}

func main() {
	fmt.Println(getRow(2))
}
