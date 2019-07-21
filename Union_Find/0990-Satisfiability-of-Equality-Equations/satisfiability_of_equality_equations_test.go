package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquationsPossible(t *testing.T) {
	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("EquationsPossible(%d)", i), func(t *testing.T) {
			assert.Equal(t, test.expected, equationsPossible(test.equations))
		})
	}
}

var result bool

func BenchmarkEquationsPossible(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r bool
			for _, test := range TestCases {
				r = equationsPossible(test.equations)
			}
			result = r
		}
	})
}
