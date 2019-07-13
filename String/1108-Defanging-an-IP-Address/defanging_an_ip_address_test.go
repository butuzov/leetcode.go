package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefangIPaddr(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("DefangIPaddr(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, defangIPaddr(test.address), "Input(address=(%+v))", test.address)
		})
	}
}

func BenchmarkDefangIPaddr(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				defangIPaddr(test.address)
			}
		}
	})
}
