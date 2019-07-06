package main
import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInterleave(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("IsInterleave(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, isInterleave(test.s1, test.s2, test.s3))
		})
	}
}

func BenchmarkIsInterleave(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				isInterleave(test.s1, test.s2, test.s3)
			}
		}
	})
}
