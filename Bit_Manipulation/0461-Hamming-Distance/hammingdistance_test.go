package main
import (
	"math"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func hammingDistanceNaive(x int, y int) int {
	var i uint
	var sum int
	for i = 0; i < uint(math.Max(float64(x), float64(y))); i++ {
		if x&(1<<i) != y&(1<<i) {
			sum++
		}
	}
	return sum
}

func hammingDistance(x int, y int) int {
	// Source - wikipedia
	var val = x ^ y
	var sum int
	for val != 0 {
		sum++
		val &= val - 1
	}
	return sum
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input X(%08b) & Y(%08b): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input X(%08b) & Y(%08b): as Expected(%+v)"

var TestCases = []struct {
	x, y     int
	expected int
}{
	{1, 4, 2},
	{680142203, 1111953568, 19},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := hammingDistance(test.x, test.y)

		if actual != test.expected {
			t.Errorf(MessageError, test.x, test.y, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.x, test.y, test.expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			hammingDistance(test.x, test.y)
		}
	}
}
