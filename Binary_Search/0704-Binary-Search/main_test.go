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

func TestSearch(t *testing.T) {
	test(t, search)
}

func BenchmarkSearch(b *testing.B) {
	bench(b, search)
}

var TestCases = []struct {
	nums     []int
	target   int
	expected int
}{
	{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
	{[]int{-1, 0, 3, 5, 9, 12}, 100, -1},
	{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
}

// Helper methods --------------------------------------------------------------

func funcName(fnc func([]int, int) int) string {
	ptr := reflect.ValueOf(fnc).Pointer()
	path := runtime.FuncForPC(ptr).Name()
	names1 := strings.Split(path, "/")
	names2 := strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func([]int, int) int) {
	name := funcName(fnc)
	for i, test := range TestCases {
		test := test
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.expected, fnc(test.nums, test.target))
		})
	}
}

// prevention of inline optimization.
var result int

func bench(b *testing.B, fnc func([]int, int) int) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r int
			for _, test := range TestCases {
				r = fnc(test.nums, test.target)
			}
			result = r
		}
	})
}
