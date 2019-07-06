package main
import (
	"math"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func countPrimesNaive(num int) int {

	var primes []int
	var isPrime = false

	for n := 2; n < num; n++ {
		isPrime = true

		for m := 2; m < int(math.Sqrt(float64(n)))+1; m++ {
			if n%m == 0 {
				isPrime = false
				break
			}
		}

		if isPrime == false {
			continue
		}
		primes = append(primes, n)

	}

	return len(primes)
}

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	table := make([]bool, n)
	table[0], table[1] = true, true
	primes := 0
	for i := 2; i < n; i++ {
		if table[i] {
			continue
		}
		primes++
		for j := i * i; j < n; j += i {
			table[j] = true
		}
	}
	return primes
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input, expected int
}{
	{10, 4},
	{1, 0},
	{2, 0},
	{3, 1},
	{4, 2},
	{999983, 78497},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := countPrimes(test.input)

		if actual != test.expected {
			t.Errorf(MessageError, test.input, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.input, test.expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			countPrimes(test.input)
		}
	}
}
