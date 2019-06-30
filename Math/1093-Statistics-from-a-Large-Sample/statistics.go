package statistics

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func sampleStats(nums []int) []float64 {
	var stats = []float64{
		-1, // min
		-1, // max
		0,  // mean
		0,  // median
		0,  // mode
	}

	var avgTotal, avgSum int

	var modeMax, modeel int

	for i, n := range nums {
		if n == 0 {
			continue
		}

		// minimum
		if stats[0] == -1 && n > 0 {
			stats[0] = float64(i)
		}

		// maximum
		stats[1] = float64(i)

		// pre_step: average/mean
		avgTotal += n
		avgSum += n * i

		// pre_step: mode
		if modeMax < n {
			modeMax = n
			modeel = i
		}

	}

	// final average calc
	stats[2] = float64(avgSum) / float64(avgTotal)

	// median
	// part 1: picking indexes fo the median elements
	var elements = []int{}
	if avgTotal%2 == 0 {
		elements = append(elements, (avgTotal/2)-1)
	}
	elements = append(elements, avgTotal/2)

	var pick = func(index int) int {
		var num int
		var offset int

		for i, n := range nums {

			if n == 0 {
				continue
			}

			if offset+n-1 < index {
				offset += n
				continue
			}

			num = i
			break

		}

		return num
	}

	// actually picking elements
	for _, idx := range elements {
		stats[3] += float64(pick(idx))
	}
	stats[3] /= float64(len(elements))

	// mode
	stats[4] = float64(modeel)

	return stats
}
