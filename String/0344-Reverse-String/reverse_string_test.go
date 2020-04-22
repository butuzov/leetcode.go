package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/butuzov/leetcode.go/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// funcName for easier name representation
func funcName(fnc func([]byte)) string {
	var ptr = reflect.ValueOf(fnc).Pointer()
	var path = runtime.FuncForPC(ptr).Name()
	var names1 = strings.Split(path, "/")
	var names2 = strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func([]byte)) {

	var name = funcName(fnc)

	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			text := utils.CopySliceByte(test.text)
			fnc(text)
			assert.Equal(t, test.expected, text)
		})
	}
}

// prevention of inline optimization.
var result []byte

func bench(b *testing.B, fnc func([]byte)) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r []byte
			for _, test := range TestCases {
				text := utils.CopySliceByte(test.text)
				fnc(text)
				r = text
			}
			result = r
		}
	})
}

// Actual Benchmarks & Tests -------------------------------------

func TestReverseString(t *testing.T) {
	test(t, reverseString)
}

func BenchmarkReverseString(b *testing.B) {
	bench(b, reverseString)
}
