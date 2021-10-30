package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// funcName for easier name representation
func funcName(fnc func([]int, []int) int) string {
	ptr := reflect.ValueOf(fnc).Pointer()
	path := runtime.FuncForPC(ptr).Name()
	names1 := strings.Split(path, "/")
	names2 := strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func([]int, []int) int) {
	name := funcName(fnc)

	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			assert.Equal(t, test.expected, fnc(test.arr1, test.arr2))
		})
	}
}

// prevention of inline optimization.
var result int

func bench(b *testing.B, fnc func([]int, []int) int) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r int
			for _, test := range TestCases {
				r = fnc(test.arr1, test.arr2)
			}
			result = r
		}
	})
}

// Actual Benchmarks & Tests -------------------------------------

func TestMaxAbsValExpr(t *testing.T) {
	test(t, maxAbsValExpr)
}

func BenchmarkMaxAbsValExpr(b *testing.B) {
	bench(b, maxAbsValExpr)
}
