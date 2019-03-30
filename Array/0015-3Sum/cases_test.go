package trisum

var TestCases = []struct {
	input    []int
	expected [][]int
}{
	{
		[]int{-2, 1, 1},
		[][]int{
			[]int{-2, 1, 1},
		},
	},
	{
		[]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0},
		[][]int{
			[]int{0, 0, 0},
			[]int{-2, 1, 1},
			[]int{-2, -2, 4},
			[]int{-4, 1, 3},
			[]int{-4, 0, 4},
			[]int{-5, 1, 4},
		},
	},
	{
		[]int{-1, 0, 1, 2, -1, -4},
		[][]int{
			[]int{-1, 0, 1},
			[]int{-1, -1, 2},
		},
	},
}
