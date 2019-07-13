package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("MyAtoi(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, myAtoi(test.str), "Input(str=(%+v))", test.str)
		})
	}
}

func BenchmarkMyAtoi(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				myAtoi(test.str)
			}
		}
	})
}
