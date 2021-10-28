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
func funcName(fnc func(string) []string) string {
	ptr := reflect.ValueOf(fnc).Pointer()
	path := runtime.FuncForPC(ptr).Name()
	names1 := strings.Split(path, "/")
	names2 := strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func(string) []string) {
	name := funcName(fnc)

	for i, test := range TestCases {
		test := test
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.expected, fnc(test.digits))
		})
	}
}

// prevention of inline optimization.
var result []string

func bench(b *testing.B, fnc func(string) []string) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r []string
			for _, test := range TestCases {
				r = fnc(test.digits)
			}
			result = r
		}
	})
}

// Actual Benchmarks & Tests -------------------------------------

func TestLetterCombinations(t *testing.T) {
	test(t, letterCombinations)
}

func BenchmarkLetterCombinations(b *testing.B) {
	bench(b, letterCombinations)
}
