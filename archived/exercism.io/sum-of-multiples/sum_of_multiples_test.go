package summultiples

import (
	"testing"
)

func TestSumMultiples(t *testing.T) {
	var n int
	for _, test := range varTests {
		s := SumMultiples(test.limit, test.divisors...)
		if s != test.sum {
			n++
			t.Errorf("% 2d Sum of multiples of %v to %d returned %d, want %d.",
				n, test.divisors, test.limit, s, test.sum)
		}
	}
}

func BenchmarkSumMultiples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range varTests {
			SumMultiples(test.limit, test.divisors...)
		}
	}
}

func BenchmarkSumMultiplesNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range varTests {
			SumMultiplesNaive(test.limit, test.divisors...)
		}
	}
}
