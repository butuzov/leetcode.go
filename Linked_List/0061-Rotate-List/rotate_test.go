package rotatelist

import (
	"testing"
)

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := rotateRight(test.In, test.K)

		if actual.String() != test.Out.String() {
			t.Errorf(MessageError, test.K, test.In, test.Out, actual)
		} else {
			t.Logf(MessageOk, test.K, test.In, test.Out)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			rotateRight(test.In, test.K)
		}
	}
}
