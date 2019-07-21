package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplexNumberMultiply(t *testing.T) {
	t.Parallel()
	for _, test := range TestCases {
		t.Run(fmt.Sprintf("Input(a=(%s),b=(%s))", test.a, test.b), func(t *testing.T) {
			assert.Equal(t, test.expected, complexNumberMultiply(test.a, test.b))
		})
	}
}

func BenchmarkComplexNumberMultiply(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				complexNumberMultiply(test.a, test.b)
			}
		}
	})
}
