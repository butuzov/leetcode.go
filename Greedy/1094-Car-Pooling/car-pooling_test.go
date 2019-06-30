package carpooling

import (
	"sort"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func carPooling(trips [][]int, capacity int) bool {

	var stops = make(map[int]int)
	for _, stop := range trips {
		stops[stop[1]] += stop[0]
		stops[stop[2]] -= stop[0]
	}

	var sortedStops = make([]int, len(stops))
	var index int
	for i := range stops {
		sortedStops[index] = i
		index++
	}
	sort.Ints(sortedStops)

	var stopIndex int
	for i := 0; i < len(sortedStops); i++ {
		stopIndex = sortedStops[i]
		capacity -= stops[stopIndex]

		if capacity < 0 {
			return false
		}
	}

	return true
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v, %+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v, %+v) as Expected(%+v)"

var TestCases = []struct {
	trips    [][]int
	capacity int
	expected bool
}{
	{
		[][]int{{2, 1, 5}, {3, 3, 7}},
		4,
		false,
	},
	{
		[][]int{{2, 1, 5}, {3, 3, 7}},
		5,
		true,
	},

	{
		[][]int{{2, 1, 5}, {3, 5, 7}},
		3,
		true,
	},
	{
		[][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}},
		11,
		true,
	},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestCarPooling(t *testing.T) {
	for _, test := range TestCases {
		actual := carPooling(test.trips, test.capacity)

		if actual != test.expected {
			t.Errorf(MessageError, test.trips, test.capacity, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.trips, test.capacity, test.expected)
		}
	}
}

func BenchmarkCarPooling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			carPooling(test.trips, test.capacity)
		}
	}
}
