package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTitleToNumber(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("TitleToNumber(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, titleToNumber(test.s), "Input(s=(%+v))", test.s)
		})
	}
}

func BenchmarkTitleToNumber(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				titleToNumber(test.s)
			}
		}
	})
}
