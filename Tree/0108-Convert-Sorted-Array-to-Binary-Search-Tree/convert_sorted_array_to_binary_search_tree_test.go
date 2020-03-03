package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/butuzov/leetcode.go/pkg/binarytree"
	"github.com/stretchr/testify/assert"
)

// funcName for easier name representation
func funcName(fnc func([]int) *binarytree.TreeNode) string {
	var ptr = reflect.ValueOf(fnc).Pointer()
	var path = runtime.FuncForPC(ptr).Name()
	var names1 = strings.Split(path, "/")
	var names2 = strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func([]int) *binarytree.TreeNode) {

	var name = funcName(fnc)

	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			t1, t2 := test.expected, fnc(test.nums)
			assert.Equal(t, t2.String(), t1.String())
		})
	}
}

// prevention of inline optimization.
var result *binarytree.TreeNode

func bench(b *testing.B, fnc func([]int) *binarytree.TreeNode) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r *binarytree.TreeNode
			for _, test := range TestCases {
				r = fnc(test.nums)
			}
			result = r
		}
	})
}

// Actual Benchmarks & Tests -------------------------------------

func TestSortedArrayToBST(t *testing.T) {
	test(t, sortedArrayToBST)
}

func BenchmarkSortedArrayToBST(b *testing.B) {
	bench(b, sortedArrayToBST)
}
