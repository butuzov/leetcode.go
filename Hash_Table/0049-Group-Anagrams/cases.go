package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	strs     []string
	expected [][]string
}{
	{
		[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
		[][]string{
			{"ate", "eat", "tea"},
			{"nat", "tan"},
			{"bat"},
		},
	},
	{
		[]string{"tao", "pit", "cam", "aid", "pro", "dog"},
		[][]string{
			{"tao"}, {"pit"}, {"cam"}, {"aid"}, {"pro"}, {"dog"},
		},
	},
	{
		[]string{"abc", "de", "fghij", "bca", "cba", "ed", "acb"},
		[][]string{
			{"abc", "bca", "cba", "acb"},
			{"de", "ed"},
			{"fghij"},
		},
	},
}
