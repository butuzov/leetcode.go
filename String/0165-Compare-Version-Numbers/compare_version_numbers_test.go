package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareVersion(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("CompareVersion(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, compareVersion(test.version1, test.version2), "Input(version1=(%s),version2=(%s))", test.version1, test.version2)
		})
	}
}

func BenchmarkCompareVersion(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				compareVersion(test.version1, test.version2)
			}
		}
	})
}
