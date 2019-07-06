package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumNaive(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("TwoSum(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, TwoSumNaive(test.nums, test.target))
		})
	}
}

func TestTwoSum(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("TwoSum(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, TwoSum(test.nums, test.target))
		})
	}
}

func benchSimple(b *testing.B, fnc func([]int, int) []int) {
	for n := 0; n < b.N; n++ {
		for _, test := range TestCases {
			fnc(test.nums, test.target)
		}
	}
}

func benchAdvanced(b *testing.B, fnc func([]int, int) []int) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				fnc(test.nums, test.target)
			}
		}
	})
}

func BenchmarkTwoSum(b *testing.B) {
	benchSimple(b, TwoSum)
}
func BenchmarkTwoSumNaive(b *testing.B) {
	benchSimple(b, TwoSumNaive)
}

func BenchmarkTwoSum_Paralel(b *testing.B) {
	benchAdvanced(b, TwoSum)
}

func BenchmarkTwoSumNaive_Paralel(b *testing.B) {
	benchAdvanced(b, TwoSumNaive)
}
