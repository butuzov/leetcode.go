package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDisappearedNumbers(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("FindDisappearedNumbers(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, findDisappearedNumbers(test.nums), "Input(nums=(%+v))", test.nums)
		})
	}
}

var result []int

func BenchmarkFindDisappearedNumbers(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r []int
			for _, test := range TestCases {
				r = findDisappearedNumbers(test.nums)
			}
			result = r
		}
	})
}
