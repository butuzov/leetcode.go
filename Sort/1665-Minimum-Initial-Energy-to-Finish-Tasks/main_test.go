package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Actual Benchmarks & Tests ---------------------------------------------------

func TestMinimumEffort(t *testing.T) {
	test(t, minimumEffort)
}

func BenchmarkMinimumEffort(b *testing.B) {
	bench(b, minimumEffort)
}

var TestCases = []struct {
	tasks    [][]int
	expected int
}{
	{
		tasks:    [][]int{{1, 3}, {2, 4}, {10, 11}, {10, 12}, {8, 9}},
		expected: 32,
	},
	{
		tasks:    [][]int{{1, 2}, {2, 4}, {4, 8}},
		expected: 8,
	},
	{
		tasks:    [][]int{{1, 7}, {2, 8}, {3, 9}, {4, 10}, {5, 11}, {6, 12}},
		expected: 27,
	},
}

// Helper methods --------------------------------------------------------------

func funcName(fnc func([][]int) int) string {
	ptr := reflect.ValueOf(fnc).Pointer()
	path := runtime.FuncForPC(ptr).Name()
	names1 := strings.Split(path, "/")
	names2 := strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func([][]int) int) {
	name := funcName(fnc)
	for i, test := range TestCases {
		test := test
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.expected, fnc(test.tasks))
		})
	}
}

// prevention of inline optimization.
var result int

func bench(b *testing.B, fnc func([][]int) int) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r int
			for _, test := range TestCases {
				r = fnc(test.tasks)
			}
			result = r
		}
	})
}
