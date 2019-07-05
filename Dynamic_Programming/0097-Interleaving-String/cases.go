package interleavingstring

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	s1       string
	s2       string
	s3       string
	expected bool
}{
	{
		"",
		"",
		"",
		true,
	},
	{
		"a",
		"b",
		"ab",
		true,
	},
	{
		"a",
		"",
		"c",
		false,
	},
	{
		"ba",
		"ab",
		"abba",
		true,
	},
	{
		"aaa",
		"aaaa",
		"aaaaaaa",
		true,
	},

	{
		"aabcc",
		"dbbca",
		"aadbbcbcac",
		true,
	},
	{
		"aabcc",
		"dbbca",
		"aadbbbaccc",
		false,
	},
	{
		"bbbbbabbbbabaababaaaabbababbaaabbabbaaabaaaaababbbababbbbbabbbbababbabaabababbbaabababababbbaaababaa",
		"babaaaabbababbbabbbbaabaabbaabbbbaabaaabaababaaaabaaabbaaabaaaabaabaabbbbbbbbbbbabaaabbababbabbabaab",
		"babbbabbbaaabbababbbbababaabbabaabaaabbbbabbbaaabbbaaaaabbbbaabbaaabababbaaaaaabababbababaababbababbbababbbbaaaabaabbabbaaaaabbabbaaaabbbaabaaabaababaababbaaabbbbbabbbbaabbabaabbbbabaaabbababbabbabbab",
		false,
	},
}
