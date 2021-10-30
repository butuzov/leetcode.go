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

func Test{{.FuncNameCapital}}(t *testing.T) {
	test(t, {{.FuncName}})
}

func Benchmark{{.FuncNameCapital}}(b *testing.B) {
	bench(b, {{.FuncName}})
}


var TestCases = []struct{
{{.TestStruct}}
}{}

// Helper methods --------------------------------------------------------------

func funcName(fnc {{.Func}}) string {
	ptr := reflect.ValueOf(fnc).Pointer()
	path := runtime.FuncForPC(ptr).Name()
	names1 := strings.Split(path, "/")
	names2 := strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc {{.Func}}) {
	name := funcName(fnc)
	for i, test := range TestCases {
		test := test
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.expected, fnc({{.TestArgs}}))
		})
	}
}

// prevention of inline optimization.
var result {{.ReturnType}}

func bench(b *testing.B, fnc {{.Func}}) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r {{.ReturnType}}
			for _, test := range TestCases {
				r = fnc({{.TestArgs}})
			}
			result = r
		}
	})
}

