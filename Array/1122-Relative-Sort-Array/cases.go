package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	arr1     []int
	arr2     []int
	expected []int
}{
	{[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}, []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}},
	{[]int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}, []int{22, 28, 8, 6, 17, 44}},
}