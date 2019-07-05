package bestsightseeingpair

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxScoreSightseeingPairt(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("MaxScoreSightseeingPair(%d)", i), func(t *testing.T) {
			assert.Equalf(t, test.expected, maxScoreSightseeingPairNaiveMomoization(test.A), "Input(A=(%+v))", test.A)
		})
	}
}

func TestMaxScoreSightseeingPairNaiveMomoization(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("MaxScoreSightseeingPair(%d)", i), func(t *testing.T) {
			assert.Equalf(t, test.expected, maxScoreSightseeingPairNaiveMomoization(test.A), "Input(A=(%+v))", test.A)
		})
	}
}

func BenchmarkMaxSScoreSightseeingPairNaiveMomoization(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				maxScoreSightseeingPairNaiveMomoization(test.A)
			}
		}
	})
}

func BenchmarkMaxScoreSightseeingPair(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				maxScoreSightseeingPair(test.A)
			}
		}
	})
}
