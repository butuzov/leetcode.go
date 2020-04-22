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
func funcName(fnc func(string) bool) string {
	var ptr = reflect.ValueOf(fnc).Pointer()
	var path = runtime.FuncForPC(ptr).Name()
	var names1 = strings.Split(path, "/")
	var names2 = strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func(string) bool) {

	var name = funcName(fnc)

	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			assert.Equal(t, test.expected, fnc(test.s))
		})
	}
}

// prevention of inline optimization.
var result bool

func bench(b *testing.B, fnc func(string) bool) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r bool
			for _, test := range TestCases {
				r = fnc(test.s)
			}
			result = r
		}
	})
}

// Actual Benchmarks & Tests -------------------------------------

func TestIsValidString(t *testing.T) {
	test(t, isValidString)
}

func BenchmarkIsValidString(b *testing.B) {
	bench(b, isValidString)
}

func TestIsValid(t *testing.T) {
	test(t, isValid)
}

func BenchmarkIsValid(b *testing.B) {
	bench(b, isValid)
}
