package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicates(t *testing.T) {
	t.Parallel()
	for _, test := range TestCases {
		t.Run(fmt.Sprintf("FindDuplicates(%+v)", test.nums), func(t *testing.T) {
			assert.Equal(t, test.expected, findDuplicates(test.nums))
		})
	}
}

var result []int

func BenchmarkFindDuplicates(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r []int
			for _, test := range TestCases {
				r = findDuplicates(test.nums)
			}
			result = r
		}
	})
}
