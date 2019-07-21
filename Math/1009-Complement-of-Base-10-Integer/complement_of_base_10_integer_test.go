package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitwiseComplement(t *testing.T) {
	t.Parallel()
	for _, test := range TestCases {
		t.Run(fmt.Sprintf("BitwiseComplement(%d)", test.N), func(t *testing.T) {
			assert.Equal(t, test.expected, bitwiseComplement(test.N))
		})
	}
}

func BenchmarkBitwiseComplement(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				bitwiseComplement(test.N)
			}
		}
	})
}
