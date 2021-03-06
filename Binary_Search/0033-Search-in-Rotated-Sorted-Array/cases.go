package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	nums     []int
	target   int
	expected int
}{
	{[]int{}, 0, -1},
	{[]int{4}, 0, -1},
	{[]int{4}, 4, 0},
	{[]int{4, 5}, 0, -1},
	{[]int{4, 5}, 4, 0},

	{[]int{4, 5, 6, 7, 0, 1, 2}, 5, 1},

	{[]int{4, 5, 6, 0, 1, 2, 3}, 1, 4},
	{[]int{4, 5, 6, 0, 1, 2, 3}, 5, 1},
	{[]int{4, 5, 6, 0, 1, 2, 3}, 4, 0},
	{[]int{4, 5, 6, 0, 1, 2, 3}, 3, 6},

	{[]int{5, 6, 7, 8, 0, 1, 2}, 3, -1},
	{[]int{5, 6, 7, 8, 0, 1, 2}, 4, -1},
	{[]int{279, 284, 285, 287, 288, 296, 2, 8, 10, 11, 12, 13, 14, 19, 20, 22, 25, 26, 27, 28, 29, 33, 37, 39, 42, 43, 47, 48, 49, 50, 57, 62, 63, 66, 68, 69, 71, 73, 74, 77, 78, 79, 85, 89, 92, 94, 99, 111, 113, 117, 120, 122, 125, 134, 145, 152, 155, 157, 160, 161, 167, 168, 181, 182, 188, 189, 190, 194, 199, 201, 204, 208, 213, 220, 223, 225, 226, 227, 231, 237, 240, 242, 247, 254, 259, 260, 261, 264, 267, 271, 275}, 298, -1},
}
