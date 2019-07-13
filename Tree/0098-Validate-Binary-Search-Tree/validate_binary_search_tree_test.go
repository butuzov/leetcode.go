package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidBSTEveryNode(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("IsValidBST(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, isValidBSTEveryNode(test.root))
		})
	}
}

func TestIsValidBST(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("IsValidBST(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, isValidBST(test.root))
		})
	}
}

func BenchmarkIsValidBSTEveryNode(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				isValidBSTEveryNode(test.root)
			}
		}
	})
}

func BenchmarkIsValidBST(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				isValidBST(test.root)
			}
		}
	})
}
