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

func TestDayOfYear(t *testing.T) {
	test(t, dayOfYear)
}

func BenchmarkDayOfYear(b *testing.B) {
	bench(b, dayOfYear)
}

var TestCases = []struct {
	date     string
	expected int
}{
	{"2019-01-09", 9},
	{"2019-02-10", 41},
	{"2003-03-01", 60},
	{"2004-03-01", 61},
}

// Helper methods --------------------------------------------------------------

func funcName(fnc func(string) int) string {
	ptr := reflect.ValueOf(fnc).Pointer()
	path := runtime.FuncForPC(ptr).Name()
	names1 := strings.Split(path, "/")
	names2 := strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func(string) int) {
	name := funcName(fnc)
	for i, test := range TestCases {
		test := test
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.expected, fnc(test.date))
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
				r = fnc(test.date)
			}
			result = r
		}
	})
}
