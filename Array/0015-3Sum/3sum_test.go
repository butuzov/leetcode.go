package trisum

import (
	"reflect"
	"sort"
	"testing"
)

//  winer solution
//  not mine
func threeSum(nums []int) [][]int {
	// change argument
	sort.Ints(nums)
	var result [][]int

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		start := i + 1
		end := len(nums) - 1

		for start < end && start < len(nums) && end > i {
			sums := nums[start] + nums[end] + nums[i]
			if sums == 0 {
				result = append(result, []int{nums[i], nums[start], nums[end]})
				for start++; start < end && nums[start] == nums[start-1]; start++ {
				}
				for end--; start < end && nums[end] == nums[end+1]; end-- {
				}
			} else if sums < 0 { // increase start till it's not equal to the previous
				start++
			} else { // decrease end till it's not equal to the previous
				end--
			}
		}
	}

	return result
}

/*******************************************************************************
  Problem Solution

  This is almost naive solution, ins't really optimal.

	Runtime: 3292 ms, faster than 9.23% of Go online submissions for 3Sum.
	Memory Usage: 279.3 MB, less than 54.55% of Go online submissions for 3Sum.
*******************************************************************************/
func threeSumNaive(nums []int) [][]int {
	var tmp [][]int

	comparator := func(a, b []int) {

		for i := 0; i < len(a); i++ {
			for j := i + 1; j < len(a); j++ {

				for n := 0; n < len(b); n++ {
					if a[i]+a[j]+b[n] != 0 {
						continue
					}

					var ints = []int{a[i], a[j], b[n]}
					sort.Ints(ints)
					tmp = append(tmp, ints)
					break
				}

			}
		}
	}

	var negs, poss = []int{}, []int{}
	for _, i := range nums {
		if i < 0 {
			negs = append(negs, i)
			continue
		}
		poss = append(poss, i)
	}

	sort.Ints(poss)
	sort.Ints(negs)

	// special case of 000
	if len(poss) >= 3 && poss[0] == 0 && poss[1] == 0 && poss[2] == 0 {
		tmp = append(tmp, []int{0, 0, 0})
	}

	comparator(poss, negs)
	comparator(negs, poss)

	var results [][]int
	var mapa = make(map[struct{ a, b, c int }]bool)

	for _, v := range tmp {
		if _, ok := mapa[struct{ a, b, c int }{v[0], v[1], v[2]}]; ok != true {
			mapa[struct{ a, b, c int }{v[0], v[1], v[2]}] = true
			results = append(results, v)
		}
	}

	return results
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {

	for _, test := range TestCases {
		actual := threeSum(test.input)

		if len(actual) != len(test.expected) {
			t.Errorf(MessageError, test.input, test.expected, actual)
			continue
		}

		sort.Slice(actual, func(i, j int) bool {
			for index := range actual[i] {
				if actual[i][index] > actual[j][index] {
					return true
				}
			}
			return false
		})

		// fmt.Println("T", test.expected)
		// fmt.Println("A", actual)

		// fmt.Println()

		for i := range actual {
			if !reflect.DeepEqual(actual[i], test.expected[i]) {
				t.Errorf(MessageError, test.input, test.expected, actual)
				break
			}
		}

	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			threeSum(test.input)
		}
	}
}
