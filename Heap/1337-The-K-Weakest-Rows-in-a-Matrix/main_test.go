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

func TestKWeakestRows(t *testing.T) {
	test(t, kWeakestRows)
}

func BenchmarkKWeakestRows(b *testing.B) {
	bench(b, kWeakestRows)
}

var TestCases = []struct {
	mat      [][]int
	k        int
	expected []int
}{
	{
		[][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 1},
		},
		3,
		[]int{2, 0, 3},
	},

	{
		[][]int{
			{1, 0, 0, 0},
			{1, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
		},
		2,
		[]int{0, 2},
	},

	{
		[][]int{
			{1, 0},
			{1, 0},
			{1, 0},
			{1, 1},
		},
		4,
		[]int{0, 1, 2, 3},
	},
}

// Helper methods --------------------------------------------------------------

func funcName(fnc func([][]int, int) []int) string {
	ptr := reflect.ValueOf(fnc).Pointer()
	path := runtime.FuncForPC(ptr).Name()
	names1 := strings.Split(path, "/")
	names2 := strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func([][]int, int) []int) {
	name := funcName(fnc)
	for i, test := range TestCases {
		test := test
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.expected, fnc(test.mat, test.k))
		})
	}
}

// prevention of inline optimization.
var result []int

func bench(b *testing.B, fnc func([][]int, int) []int) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r []int
			for _, test := range TestCases {
				r = fnc(test.mat, test.k)
			}
			result = r
		}
	})
}
