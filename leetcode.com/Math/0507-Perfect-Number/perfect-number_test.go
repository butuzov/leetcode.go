package perfectnumber

import (
	"math"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func checkPerfectNumber(num int) bool {

	if num < 6 {
		return false
	}

	var cmp = num
	var sum = 1
	var divisors = map[int]int{
		1: num,
	}

	var increment = 1
	if num%2 == 1 {
		increment = 2
	}

	for i := 1; i < num; i += increment {

		if _, ok := divisors[i]; ok == true {
			continue
		}

		divisor, modulo := cmp/i, cmp%i

		if modulo != 0 {
			continue
		}

		divisors[i] = divisor
		divisors[divisor] = i

		num = divisor

		sum += i
		sum += divisor
	}

	if sum == cmp {
		return true
	}

	return false
}

// This code by unknown author provides way better
// performance
func checkPerfectNumberSimple(num int) bool {
	if num < 6 {
		return false
	}

	tmp := num - 1
	for i := 2; i < int(math.Ceil(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			tmp -= i + (num / i)
		}
		if tmp < 0 {
			return false
		}
	}

	if tmp == 0 {
		return true
	}
	return false
}

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

var TestCases = []struct {
	input    int
	expected bool
}{
	{28, true},
	{6, true},
	{0, false},
	{1, false},
	{99999997, false},
	{100000000, false},
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestProblem(t *testing.T) {
	for _, test := range TestCases {
		actual := checkPerfectNumber(test.input)

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
			checkPerfectNumber(test.input)
		}
	}
}

func BenchmarkProblemSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			checkPerfectNumberSimple(test.input)
		}
	}
}
