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

func TestMaxPerformance(t *testing.T) {
	test(t, maxPerformance)
}

func BenchmarkMaxPerformance(b *testing.B) {
	bench(b, maxPerformance)
}

var TestCases = []struct {
	n          int
	speed      []int
	efficiency []int
	k          int
	expected   int
}{

	{
		n:          6,
		speed:      []int{2, 10, 3, 1, 5, 8},
		efficiency: []int{5, 4, 3, 9, 7, 2},
		k:          2,
		expected:   60, // We have the maximum performance of the team by selecting engineer 2 (with speed=10 and efficiency=4) and engineer 5 (with speed=5 and efficiency=7). That is, performance = (10 + 5) * min(4, 7) = 60.
	},

	{
		n:          6,
		speed:      []int{2, 10, 3, 1, 5, 8},
		efficiency: []int{5, 4, 3, 9, 7, 2},
		k:          3,
		expected:   68, // his is the same example as the first but k = 3. We can select engineer 1, engineer 2 and engineer 5 to get the maximum performance of the team. That is, performance = (2 + 10 + 5) * min(5, 4, 7) = 68.
	},

	{
		n:          6,
		speed:      []int{2, 10, 3, 1, 5, 8},
		efficiency: []int{5, 4, 3, 9, 7, 2},
		k:          4,
		expected:   72,
	},
}

// Helper methods --------------------------------------------------------------

func funcName(fnc func(int, []int, []int, int) int) string {
	ptr := reflect.ValueOf(fnc).Pointer()
	path := runtime.FuncForPC(ptr).Name()
	names1 := strings.Split(path, "/")
	names2 := strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func(int, []int, []int, int) int) {
	name := funcName(fnc)
	for i, test := range TestCases {
		test := test
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.expected, fnc(test.n, test.speed, test.efficiency, test.k))
		})
	}
}

// prevention of inline optimization.
var result int

func bench(b *testing.B, fnc func(int, []int, []int, int) int) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r int
			for _, test := range TestCases {
				r = fnc(test.n, test.speed, test.efficiency, test.k)
			}
			result = r
		}
	})
}
