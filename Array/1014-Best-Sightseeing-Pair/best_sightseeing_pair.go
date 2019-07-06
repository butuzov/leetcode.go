package main
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*******************************************************************************
  Problem Naive Solution

*******************************************************************************/

type doublet struct {
	a, b int
}

func maxScoreSightseeingPairNaiveMomoization(A []int) int {

	var cache = make(map[doublet]int)

	var result = func(i, j int) int {

		var (
			val int
			ok  bool
			key = doublet{i, j}
		)

		if j > i {
			key.a = j
			key.b = i
		}

		if val, ok = cache[key]; !ok {
			val = i + j
			cache[key] = val
		}
		return val
	}

	var target = -1
	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < len(A); j++ {
			target = max(target, result(A[i], A[j])-(j-i))
		}
	}

	return target
}

// --- comunity solution.

func maxScoreSightseeingPair(A []int) int {

	var maxScore, locScore int
	maxScore = A[0]

	for i := 0; i < len(A); i++ {
		maxScore = max(maxScore, locScore+A[i]-i)
		locScore = max(locScore, A[i]+i)
	}
	return maxScore
}
