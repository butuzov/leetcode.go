package pascal

import "fmt"

// This example use two implementation - one with factorials and other one with
// simple sum for previous row. One with factorials allow ot calculate random
// row of pascal's triangle.

// Please NOTE - this factorial will work with default int.
// we need to specify int64 in order to work with bigger numbers.

// memoizeFunction declares way we will use it with factorial function.
type memoizeFunction func(int) interface{}

// declaring globaly factorial
var factorial memoizeFunction

// Simple decorator we use to speedup factorils.
func memoize(function memoizeFunction) memoizeFunction {

	mapa := make(map[int]interface{})

	return func(n int) interface{} {

		if value, ok := mapa[n]; ok {
			return value
		}

		mapa[n] = function(n)
		return mapa[n]
	}
}

// We use init to declar factorial callbackm that uses memoize
func init() {
	factorial = memoize(func(n int) interface{} {
		if n == 0 {
			return 1
		}
		return factorial(n-1).(int) * n
	})
}

// Triangle implements way to create Pascal's triangle by
// createng each leavel separatly.
func Triangle(height int) [][]int {
	var triangle [][]int
	triangle = make([][]int, height, height)
	for n := 1; n <= height; n++ {
		triangle[n-1] = TriangleRow(n)
	}
	return triangle
}

// TriangleRow - creates a level for Pascals triangle.
// It uses combinatorics funcion with meomoization technic
// to spead up calculatons. Factorial function defined in init method.
// its possible to optimize even more by simplifing factorials, but it will
// make code more hard to understand (it require some math background
// in combinatorics).
func TriangleRow(n int) []int {
	var row = make([]int, n)

	var center, rest = n / 2, n % 2
	center += rest

	// create right part
	for i := 0; i < center; i++ {
		row[i] = factorial(n-1).(int) / (factorial(n-i-1).(int) * factorial(i).(int))
	}

	// just copy left part
	for i := n - 1; i >= center; i-- {
		row[i] = row[n-i-1]
	}

	return row
}

// TriangleNaive implements simple and even faster way
// to calculate Pascals triangle, but based on previous row.
func TriangleNaive(height int) [][]int {
	var triangle [][]int
	triangle = make([][]int, height, height)
	for n := 1; n <= height; n++ {
		var row = n - 1
		triangle[row] = make([]int, n)

		// define center of the row
		var center, rest = n / 2, n % 2
		center += rest

		triangle[row][0] = 1

		for i := 1; i < center; i++ {
			triangle[row][i] = triangle[row-1][i-1] + triangle[row-1][i]
		}

		// just copy left part
		for i := n - 1; i >= center; i-- {
			triangle[row][i] = triangle[row][n-i-1]
		}

	}
	return triangle
}

func main() {
	// var limit int
	// limit = 9
	// fmt.Printf("%d\n", factorial(10))
	for i, row := range TriangleNaive(7) {
		fmt.Printf("%d] %v\n", i, row)
	}

}
