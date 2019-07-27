package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	A        []string
	B        []string
	expected []string
}{
	{
		[]string{"amazon", "apple", "facebook", "google", "leetcode"},
		[]string{"e", "o"},
		[]string{"facebook", "google", "leetcode"},
	},
	{
		[]string{"amazon", "apple", "facebook", "google", "leetcode"},
		[]string{"l", "e"},
		[]string{"apple", "google", "leetcode"},
	},
	{
		[]string{"amazon", "apple", "facebook", "google", "leetcode"},
		[]string{"e", "oo"},
		[]string{"facebook", "google"},
	},

	{
		[]string{"amazon", "apple", "facebook", "google", "leetcode"},
		[]string{"lo", "eo"},
		[]string{"google", "leetcode"},
	},

	{
		[]string{"amazon", "apple", "facebook", "google", "leetcode"},
		[]string{"ec", "oc", "ceo"},
		[]string{"facebook", "leetcode"},
	},
}
