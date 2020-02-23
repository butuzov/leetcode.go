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
func funcName(fnc func([][]int) int) string {
	var ptr = reflect.ValueOf(fnc).Pointer()
	var path = runtime.FuncForPC(ptr).Name()
	var names1 = strings.Split(path, "/")
	var names2 = strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func([][]int) int) {

	var name = funcName(fnc)

	t.Parallel()
	for i := 0; i < len(TestCases); i++ {
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			assert.Equal(t, TestCases[i].expected, fnc(utils.CopyMatrix(TestCases[i].grid)))
		})
	}

}

// prevention of inline optimization.
var result int

func bench(b *testing.B, fnc func([][]int) int) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r int
			for _, test := range TestCases {
				m := utils.CopyMatrix(test.grid)
				r = fnc(m)
			}
			result = r
		}
	})
}

// Actual Benchmarks & Tests -------------------------------------

func TestMinPathSumDP(t *testing.T) {
	test(t, minPathSumDP)
}

func BenchmarkMinPathSumDP(b *testing.B) {
	bench(b, minPathSumDP)
}

func TestMinPathSumFast(t *testing.T) {
	test(t, minPathSumFast)
}

func BenchmarkMinPathSumFast(b *testing.B) {
	bench(b, minPathSumFast)
}
