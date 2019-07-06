package main
import (
	"container/heap"
	"reflect"
	"sort"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/

/*
	Heap-ed solution.
*/

type Point struct {
	p      []int
	center int
}

type Heap []*Point

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].center < h[j].center }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(*Point)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func kClosest(p [][]int, K int) [][]int {

	var results = make([][]int, 0, K)

	if K == 0 {
		return results
	}

	var pow = func() func(s int) int {
		powers := make(map[int]int)
		return func(i int) int {
			if _, ok := powers[i]; ok != true {
				powers[i] = i * i
			}
			return powers[i]
		}
	}()

	var lookUpHeap = make(Heap, len(p))

	for i, point := range p {
		square := pow(point[0]) + pow(point[1])
		lookUpHeap[i] = &Point{p: point, center: square}
	}

	heap.Init(&lookUpHeap)

	for lookUpHeap.Len() > 0 {
		item := heap.Pop(&lookUpHeap).(*Point)
		results = append(results, item.p)
		if len(results) == K {
			return results
		}
	}
	return results
}

/*
	Slice based solution.
*/
func kClosestSlice(p [][]int, K int) [][]int {

	sort.Slice(p, func(a, b int) bool {
		return p[a][0]*p[a][0]+p[a][1]*p[a][1] < p[b][0]*p[b][0]+p[b][1]*p[b][1]
	})

	var results = make([][]int, 0, K)
	for i := 0; i < K && i < len(p); i++ {
		results = append(results, p[i])
	}

	return results
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    [][]int
	kClosest int
	expected [][]int
}{
	{[][]int{[]int{-2, 2}, []int{1, 3}}, 1, [][]int{[]int{-2, 2}}},
	{[][]int{[]int{5, -1}, []int{-2, 4}, []int{3, 3}}, 2, [][]int{[]int{3, 3}, []int{-2, 4}}},
	{[][]int{[]int{5, -1}, []int{-2, 4}, []int{3, 3}}, 0, [][]int{}},
	{[][]int{[]int{5, -1}}, 2, [][]int{[]int{5, -1}}},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := kClosest(test.input, test.kClosest)

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf(MessageError, test.input, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.input, test.expected)
		}
	}
}

func TestProblemArray(t *testing.T) {
	for _, test := range TestCases {
		actual := kClosestSlice(test.input, test.kClosest)

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf(MessageError, test.input, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.input, test.expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			kClosest(test.input, test.kClosest)
		}
	}
}

func BenchmarkProblemSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			kClosestSlice(test.input, test.kClosest)
		}
	}
}
