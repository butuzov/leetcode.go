package main
import (
	"testing"
)

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := arrangeCoins2(test.input)

		if actual != test.expected {
			t.Errorf(MessageError, test.input, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.input, test.expected)
		}
	}
}

func BenchmarkProblem1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			arrangeCoins1(test.input)
		}
	}
}

func BenchmarkProblem2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			arrangeCoins2(test.input)
		}
	}
}
