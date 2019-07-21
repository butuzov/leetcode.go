package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativeSortArray(t *testing.T) {
	t.Parallel()

	for i, test := range TestCases {
		t.Run(fmt.Sprintf("RelativeSortArray(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, relativeSortArray(test.arr1, test.arr2), "Input(arr1=(%+v),=(%+v))", test.arr1, test.arr2)
		})
	}
}

func BenchmarkRelativeSortArray(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				relativeSortArray(test.arr1, test.arr2)
			}
		}
	})
}
