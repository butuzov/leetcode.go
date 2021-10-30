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

func TestGenerateParenthesis(t *testing.T) {
	test(t, generateParenthesis)
}

func BenchmarkGenerateParenthesis(b *testing.B) {
	bench(b, generateParenthesis)
}

var TestCases = []struct {
	n        int
	expected []string
}{

	{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	{1, []string{"()"}},
}

// Helper methods --------------------------------------------------------------

func funcName(fnc func(int) []string) string {
	ptr := reflect.ValueOf(fnc).Pointer()
	path := runtime.FuncForPC(ptr).Name()
	names1 := strings.Split(path, "/")
	names2 := strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func(int) []string) {
	name := funcName(fnc)
	for i, test := range TestCases {
		test := test
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.expected, fnc(test.n))
		})
	}
}

// prevention of inline optimization.
var result []string

func bench(b *testing.B, fnc func(int) []string) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r []string
			for _, test := range TestCases {
				r = fnc(test.n)
			}
			result = r
		}
	})
}
