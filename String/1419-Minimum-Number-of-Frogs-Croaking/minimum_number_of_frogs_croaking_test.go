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
func funcName(fnc func(string) int) string {
	var ptr = reflect.ValueOf(fnc).Pointer()
	var path = runtime.FuncForPC(ptr).Name()
	var names1 = strings.Split(path, "/")
	var names2 = strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func(string) int) {

	var name = funcName(fnc)

	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			assert.Equal(t, test.expected, fnc(test.croakOfFrogs), "oops %+v", test.croakOfFrogs)
		})
	}
}

// prevention of inline optimization.
var result int

func bench(b *testing.B, fnc func(string) int) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r int
			for _, test := range TestCases {
				r = fnc(test.croakOfFrogs)
			}
			result = r
		}
	})
}

// Actual Benchmarks & Tests -------------------------------------

func TestMinNumberOfFrogs(t *testing.T) {
	test(t, minNumberOfFrogs)
}

func BenchmarkMinNumberOfFrogs(b *testing.B) {
	bench(b, minNumberOfFrogs)
}
