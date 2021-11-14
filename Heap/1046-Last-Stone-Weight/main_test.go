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

func TestLastStoneWeight(t *testing.T) {
	test(t, lastStoneWeight)
}

func BenchmarkLastStoneWeight(b *testing.B) {
	bench(b, lastStoneWeight)
}


var TestCases = []struct{
	stones []int
	expected int
}{}

// Helper methods --------------------------------------------------------------

func funcName(fnc func([]int) int) string {
	ptr := reflect.ValueOf(fnc).Pointer()
	path := runtime.FuncForPC(ptr).Name()
	names1 := strings.Split(path, "/")
	names2 := strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func([]int) int) {
	name := funcName(fnc)
	for i, test := range TestCases {
		test := test
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.expected, fnc(test.stones))
		})
	}
}

// prevention of inline optimization.
var result int

func bench(b *testing.B, fnc func([]int) int) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r int
			for _, test := range TestCases {
				r = fnc(test.stones)
			}
			result = r
		}
	})
}

