package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToTitle(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("ConvertToTitle(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, convertToTitle(test.n), "Input(n=(%+v))", test.n)
		})
	}
}

func BenchmarkConvertToTitle(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				convertToTitle(test.n)
			}
		}
	})
}
