package main
import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func intToRoman(num int) string {
	if num <= 0 || num > 3999 {
		panic("Not supported")
	}

	var digits = map[int]byte{
		1000: 'M',
		500:  'D',
		100:  'C',
		50:   'L',
		10:   'X',
		5:    'V',
		1:    'I',
	}

	var buffer bytes.Buffer

	for _, grade := range []int{1000, 100, 10, 1} {
		n := int(num / grade)
		num = num % grade

		highest, _ := digits[10*grade]
		lowest, _ := digits[grade]
		middle, _ := digits[5*grade]

		switch {
		case n == 9:
			buffer.WriteByte(lowest)
			buffer.WriteByte(highest)
		case n >= 5:
			buffer.WriteByte(middle)
			buffer.WriteString(strings.Repeat(string(lowest), n-5))
		case n == 4:
			buffer.WriteByte(lowest)
			buffer.WriteByte(middle)
		default:
			buffer.WriteString(strings.Repeat(string(lowest), n))
		}
	}

	return buffer.String()
}

/*******************************************************************************
  Test Table
*******************************************************************************/

var MessageError = "Fail: In(%+v) - Expected(%v) vs Returend(%v)"
var MessageOk = "OK: In(%+v) - Return as Expected(%v)"

/*******************************************************************************
  Tests
*******************************************************************************/

func getTestTable() map[string]int {
	file, err := os.Open("romans.text")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var TestCases = make(map[string]int)

	for scanner.Scan() {
		var s = strings.Split(scanner.Text(), "	")
		i, err := strconv.Atoi(s[0])
		if err == nil {
			TestCases[s[1]] = i
		}
	}
	return TestCases
}

func TestProblem(t *testing.T) {

	for expected, input := range getTestTable() {
		actual := intToRoman(input)
		if actual != expected {
			t.Errorf(MessageError, input, expected, actual)
		} else {
			t.Logf(MessageOk, input, expected)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, input := range getTestTable() {
			intToRoman(input)
		}
	}
}
