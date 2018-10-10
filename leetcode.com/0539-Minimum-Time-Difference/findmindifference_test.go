package findmindifference

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func findMinDifference(timePoints []string) int {

	// hour to minutes
	var timenum = func(in string) int {
		parts := strings.Split(in, ":")
		h, _ := strconv.ParseInt(parts[0], 10, 0)
		m, _ := strconv.ParseInt(parts[1], 10, 0)
		return int(h)*60 + int(m)
	}

	// return absvalue of the number
	var abs = func(n int) int {
		return int(math.Abs(float64(n)))
	}

	// convert  times to number to something simplier like numbers
	times := []int{}
	for _, time := range timePoints {
		time := timenum(time)
		if time == 0 {
			times = append(times, 24*60)
		}
		times = append(times, time)
	}

	// Sorting times
	sort.Ints(times)

	// checking for min...
	var min = 24 * 60
	for i := 1; i < len(times); i++ {
		cur := abs(times[i] - times[i-1])
		if cur < min {
			min = cur
		}
	}

	// one more case last time - first time
	if times[len(times)-1] > 0 && times[len(times)-1] < 1400 {
		var nightTransition = 24*60 - times[len(times)-1] + times[0]
		if nightTransition < min {
			min = nightTransition
		}
	}
	return min
}

/*******************************************************************************
  Test Table
*******************************************************************************/
var MessageError = "Fail: Times[%+v] : Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Times[%+v]: as Expected(%+v)"

var TestCases = []struct {
	times   []string
	exected int
}{
	{[]string{"23:59", "00:00"}, 1},
	{[]string{"23:00", "00:00"}, 60},
	{[]string{"10:20", "11:10"}, 50},
	{[]string{"10:20", "11:10", "12:10", "12:01", "14:01", "14:03"}, 2},
	{[]string{"00:00", "12:00"}, 60 * 12},
	{[]string{"05:31", "22:08", "00:35"}, 147},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		result := findMinDifference(test.times)

		if test.exected != result {
			t.Errorf(MessageError, test.times, test.exected, result)
		} else {
			t.Logf(MessageOk, test.times, test.exected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			findMinDifference(test.times)
		}
	}
}
