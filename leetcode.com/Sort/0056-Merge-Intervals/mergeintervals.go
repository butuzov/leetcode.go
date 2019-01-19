package mergeintervals

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {

	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a].Start <= intervals[b].Start
	})

	var result = make([]Interval, 0, len(intervals))
	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {

		// Case 1: No Overlaps
		if intervals[i].Start > result[len(result)-1].End {
			result = append(result, intervals[i])
			continue
		}

		// Case 2: Closing Element Not In Previous Pair
		if intervals[i].End <= result[len(result)-1].End {
			continue
		}

		// Case 3: Leading Element In Previous Pair
		result[len(result)-1].End = intervals[i].End
	}

	return result
}
