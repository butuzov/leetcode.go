package main
type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, interval Interval) []Interval {
	// return empty if its empty or almost empty
	if len(intervals) < 1 {
		return append(intervals, interval)
	}

	// Results interval Slice
	var results = make([]Interval, 0, len(intervals))

	// We ned to check somehow is our interval going to be last inserted
	// we introducing a bool var we can check after the loop to insert
	// our merged interval
	var isLastInsertedInterval bool

	var tmpInterval = interval

	for _, curInterval := range intervals {

		// Pre Interval insertion
		if curInterval.End < interval.Start {
			results = append(results, curInterval)
			continue
		}

		if interval.End < curInterval.Start {

			// Merged interval insertion
			if !isLastInsertedInterval {
				results = append(results, tmpInterval)
				isLastInsertedInterval = true
			}

			// Post Interval
			results = append(results, curInterval)
			continue
		}

		// Merge interval creation
		if tmpInterval.Start == interval.Start {
			if curInterval.Start < interval.Start {
				tmpInterval.Start = curInterval.Start
			} else {
				tmpInterval.Start = interval.Start
			}
		}

		if curInterval.End > interval.End {
			tmpInterval.End = curInterval.End
		} else {
			tmpInterval.End = interval.End
		}

	}

	// Merged interval insertion
	if !isLastInsertedInterval {
		results = append(results, tmpInterval)
	}

	return results
}
