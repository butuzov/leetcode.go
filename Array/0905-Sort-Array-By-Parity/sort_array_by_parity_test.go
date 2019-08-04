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
func funcName(fnc func([]int) []int) string {
	var ptr = reflect.ValueOf(fnc).Pointer()
	var path = runtime.FuncForPC(ptr).Name()
	var names1 = strings.Split(path, "/")
	var names2 = strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func([]int) []int) {

	var name = funcName(fnc)

	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			testcase := make([]int, len(test.A))
			copy(testcase, test.A)
			assert.True(t, validate(fnc(testcase)))
		})
	}
}

// prevention of inline optimization.
var result []int

func bench(b *testing.B, fnc func([]int) []int) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r []int
			for _, test := range TestCases {

				b.StopTimer()
				testcase := make([]int, len(test.A))
				copy(testcase, test.A)
				b.StartTimer()

				r = fnc(testcase)
			}
			result = r
		}
	})
}

// Actual Benchmarks & Tests -------------------------------------

func TestSortArrayByParity(t *testing.T) {
	test(t, sortArrayByParity)
}

func BenchmarkSortArrayByParity(b *testing.B) {
	bench(b, sortArrayByParity)
}
