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

func TestFindKthLargest(t *testing.T) {
	test(t, findKthLargest)
}

func BenchmarkFindKthLargest(b *testing.B) {
	bench(b, findKthLargest)
}

func TestFindKthLargestSort(t *testing.T) {
	test(t, findKthLargestSort)
}

func BenchmarkFindKthLargestSort(b *testing.B) {
	bench(b, findKthLargestSort)
}

var TestCases = []struct {
	nums     []int
	k        int
	expected int
}{
	{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
	{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
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
			assert.Equal(t, test.expected, fnc(test.nums, test.k))
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
				r = fnc(test.nums, test.k)
			}
			result = r
		}
	})
}
