package diffsquares

// SquareOfSums return square of sum of the number
func SquareOfSums(n int) int {
	var sum int

	for i := 0; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares return sum of squares of the number
func SumOfSquares(n int) int {
	var sum int

	for i := 0; i <= n; i++ {
		sum += i * i
	}

	return sum
}

// Difference return difference of SquareOfSums and SumOfSquares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
