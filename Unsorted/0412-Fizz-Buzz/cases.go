package fizzbuzz

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	n        int
	expected []string
}{
	{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
}
