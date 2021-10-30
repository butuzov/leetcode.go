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

func TestSolveSudoku(t *testing.T) {
	test(t, solveSudoku)
}

func BenchmarkSolveSudoku(b *testing.B) {
	bench(b, solveSudoku)
}

var TestCases = []struct {
	board    [][]byte
	expected [][]byte
}{

	{
		board: [][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
		expected: [][]byte{
			{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
			{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
			{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
			{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
			{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
			{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
			{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
			{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
			{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
		},
	},
}

// Helper methods --------------------------------------------------------------

func funcName(fnc func([][]byte)) string {
	ptr := reflect.ValueOf(fnc).Pointer()
	path := runtime.FuncForPC(ptr).Name()
	names1 := strings.Split(path, "/")
	names2 := strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func([][]byte)) {
	name := funcName(fnc)
	for i, test := range TestCases {
		test := test
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			t.Parallel()
			fnc(test.board)
			assert.Equal(t, test.expected, test.expected)
		})
	}
}

// prevention of inline optimization.
func bench(b *testing.B, fnc func([][]byte)) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range TestCases {
				fnc(test.board)
			}
		}
	})
}
