package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Actual Tests ---------------------------------------------------

var TestCases = []struct {
	k        int
	nums     []int
	expected [][2]int
}{
	{
		3,
		[]int{4, 5, 8, 2},
		[][2]int{{3, 4}, {5, 5}, {10, 5}, {9, 8}, {4, 8}},
	},
}

// Helper methods --------------------------------------------------------------

func Test_KthLargest(t *testing.T) {
	for i, test := range TestCases {
		test := test
		t.Run(fmt.Sprintf("KthLargest(%d)", i), func(t *testing.T) {
			t.Parallel()

			obj := Constructor(test.k, test.nums)
			for _, pair := range test.expected {
				assert.Equal(t, pair[1], obj.Add(pair[0]))
			}
		})
	}
}
