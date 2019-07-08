package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreorderTraversal(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("PreorderTraversal(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, preorderTraversal(test.root))
		})
	}
}

func TestPreorderTraversalRecursive(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("PreorderTraversal(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, preorderTraversalRecursive(test.root))
		})
	}
}

func BenchmarkPreorderTraversal(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := range TestCases {
				preorderTraversal(TestCases[i].root)
			}
		}
	})
}

func BenchmarkPreorderTraversalRecursive(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := range TestCases {
				preorderTraversalRecursive(TestCases[i].root)
			}
		}
	})
}
