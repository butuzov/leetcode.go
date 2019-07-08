package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostorderTraversalRecursive(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("PostorderTraversal(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, postorderTraversalRecursive(test.root))
		})
	}
}

func TestPostorderTraversal(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("PostorderTraversal(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, postorderTraversal(test.root))
		})
	}
}

func BenchmarkPostorderTraversalRecursive(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := range TestCases {
				postorderTraversalRecursive(TestCases[i].root)
			}
		}
	})
}

func BenchmarkPostorderTraversal(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := range TestCases {
				postorderTraversal(TestCases[i].root)
			}
		}
	})
}
