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

func TestDaysBetweenDates(t *testing.T) {
	test(t, daysBetweenDates)
}

func BenchmarkDaysBetweenDates(b *testing.B) {
	bench(b, daysBetweenDates)
}

var TestCases = []struct {
	date1    string
	date2    string
	expected int
}{
	{
		"2019-06-29",
		"2019-06-30",
		1,
	},
	{
		"2020-01-15",
		"2019-12-31",
		15,
	},
}

// Helper methods --------------------------------------------------------------

func funcName(fnc func(string, string) int) string {
	ptr := reflect.ValueOf(fnc).Pointer()
	path := runtime.FuncForPC(ptr).Name()
	names1 := strings.Split(path, "/")
	names2 := strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func(string, string) int) {
	name := funcName(fnc)
	for i, test := range TestCases {
		test := test
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.expected, fnc(test.date1, test.date2))
		})
	}
}

// prevention of inline optimization.
var result int

func bench(b *testing.B, fnc func(string, string) int) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r int
			for _, test := range TestCases {
				r = fnc(test.date1, test.date2)
			}
			result = r
		}
	})
}
