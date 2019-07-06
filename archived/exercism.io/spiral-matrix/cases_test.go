package spiralmatrix

var testCases = []struct {
	description string
	input       int
	expected    [][]int
}{
	{
		description: "empty spiral",
		input:       0,
		expected:    [][]int{},
	},
	{
		description: "trivial spiral",
		input:       1,
		expected: [][]int{
			{1},
		},
	},
	{
		description: "spiral of size 2",
		input:       2,
		expected: [][]int{
			{1, 2},
			{4, 3},
		},
	},
	{
		description: "spiral of size 3",
		input:       3,
		expected: [][]int{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		},
	},
	{
		description: "spiral of size 4",
		input:       4,
		expected: [][]int{
			{1, 2, 3, 4},
			{12, 13, 14, 5},
			{11, 16, 15, 6},
			{10, 9, 8, 7},
		},
	},
}
