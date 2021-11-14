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

func TestFindRelativeRanks(t *testing.T) {
	test(t, findRelativeRanks)
}

func BenchmarkFindRelativeRanks(b *testing.B) {
	bench(b, findRelativeRanks)
}

var TestCases = []struct {
	score    []int
	expected []string
}{
	{
		[]int{5, 4, 3, 2, 1},
		[]string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
	},
	{
		[]int{10, 3, 8, 9, 4},
		[]string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
	},
}

// Helper methods --------------------------------------------------------------

func funcName(fnc func([]int) []string) string {
	ptr := reflect.ValueOf(fnc).Pointer()
	path := runtime.FuncForPC(ptr).Name()
	names1 := strings.Split(path, "/")
	names2 := strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func([]int) []string) {
	name := funcName(fnc)
	for i, test := range TestCases {
		test := test
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.expected, fnc(test.score))
		})
	}
}

// prevention of inline optimization.
var result []string

func bench(b *testing.B, fnc func([]int) []string) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r []string
			for _, test := range TestCases {
				r = fnc(test.score)
			}
			result = r
		}
	})
}
