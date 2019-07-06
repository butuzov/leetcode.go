package main
import "testing"

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func rob(loot []int) int {

	if len(loot) == 1 {
		return loot[0]
	}

	var max func(int, int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var cache = make(map[int]map[int]int)

	// Recursive Calls
	var lootCalc func(startIndex, endIndex int) int
	lootCalc = func(startIndex, endIndex int) int {

		if endIndex < startIndex {
			return 0
		}

		// memoizing resultes
		if starter, ok := cache[startIndex]; ok {
			if value, ok := starter[endIndex]; ok {
				return value
			}
		} else {
			cache[startIndex] = make(map[int]int)
		}

		cache[startIndex][endIndex] = loot[startIndex] + max(
			lootCalc(startIndex+2, endIndex),
			lootCalc(startIndex+3, endIndex),
		)

		return cache[startIndex][endIndex]
	}

	return max(lootCalc(0, len(loot)-1), lootCalc(1, len(loot)-1))
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := rob(test.input)
		if actual != test.expected {
			t.Errorf(MessageError, test.input, test.expected, actual)
		} else {
			t.Logf(MessageOk, test.input, test.expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			rob(test.input)
		}
	}
}
