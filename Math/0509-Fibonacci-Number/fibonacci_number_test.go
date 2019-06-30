package fib

import (
	"fmt"
	"testing"
)

var MsgError = "Fail: Input(N=(%d)): Expected(%d) vs Returend(%d)"
var MsgOk = "OK: Input(N=(%d)): as Expected(%d)"

func testRunner(t *testing.T, name string, fnc func(int) int) {
	t.Parallel()
	for i, test := range TestCases {
		test := test
		t.Run(fmt.Sprintf("%s (%d)", name, i), func(t *testing.T) {
			res := fnc(test.N)
			if res == test.expected {
				t.Logf(MsgOk, test.N, test.expected)
			} else {
				t.Errorf(MsgError, test.N, test.expected, res)
			}
		})

	}
}

func TestFibonacci(t *testing.T) {
	testRunner(t, "Fibonachi", fibonacci)
}
func TestBinets(t *testing.T) {
	testRunner(t, "Fibonachi Binet's Formula", fibonacciBinet)
}

func TestMemoize(t *testing.T) {
	testRunner(t, "Fibonachi Memoization", fibonacciMemoize)
}

func TestSwaps(t *testing.T) {
	testRunner(t, "Fibonachi Swaping fib steps", fibonacciSwaps)
}

func benchmarkRunner(fnc func(int) int) {
	for _, test := range TestCases {
		fnc(test.N)
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkRunner(fibonacci)
	}
}

func BenchmarkBinet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkRunner(fibonacciBinet)
	}
}

func BenchmarkMemoize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkRunner(fibonacciMemoize)
	}
}

func BenchmarkSwaps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkRunner(fibonacciSwaps)
	}
}
