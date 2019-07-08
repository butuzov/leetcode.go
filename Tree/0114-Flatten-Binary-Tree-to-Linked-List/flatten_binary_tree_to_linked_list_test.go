package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("Flatten(%d)", i), func(t *testing.T) {
			flatten(test)
			assert.Equal(t, test.InOrder(), test.PreOrder())
		})
	}
}

func BenchmarkFlatten(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				flatten(test)
			}
		}
	})
}
