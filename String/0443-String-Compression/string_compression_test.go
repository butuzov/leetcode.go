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
func funcName(fnc func([]byte) int) string {
	var ptr = reflect.ValueOf(fnc).Pointer()
	var path = runtime.FuncForPC(ptr).Name()
	var names1 = strings.Split(path, "/")
	var names2 = strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func([]byte) int) {

	var name = funcName(fnc)

	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			test := test
			assert.Equal(t, test.len, fnc(test.chars))
			fmt.Printf("Expected: %c\n", test.expected)
			fmt.Printf("Got:      %c\n", test.chars[0:test.len])
			assert.Equal(t, test.expected, test.chars[0:test.len])
		})
	}
}

// prevention of inline optimization.
var result int

func bench(b *testing.B, fnc func([]byte) int) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r int
			for _, test := range TestCases {
				r = fnc(test.chars)
			}
			result = r
		}
	})
}

// Actual Benchmarks & Tests -------------------------------------

func TestCompress(t *testing.T) {
	test(t, compress)
}

func BenchmarkCompress(b *testing.B) {
	bench(b, compress)
}
