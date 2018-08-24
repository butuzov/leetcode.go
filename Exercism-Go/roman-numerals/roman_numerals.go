package romannumerals

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

var numerals map[int]byte = map[int]byte{
	1000: 'M',
	500:  'D',
	100:  'C',
	50:   'L',
	10:   'X',
	5:    'V',
	1:    'I',
}

func main() {
	numeral, err := ToRomanNumeral(10)
	if err != nil {
		panic("oops")
	}
	fmt.Println(numeral)
}

// Convert Arabic numeral to Roman numeral
func ToRomanNumeral(num int) (string, error) {
	var buffer bytes.Buffer

	// roman numerals up to 3999 not supported by
	// exercism tests
	// if num <= 0 || num > 4000 {

	if num <= 0 || num > 3000 {
		return buffer.String(), errors.New("Out of roman numerals range")
	}

	var n int
	var highest, lowest, middle byte

	for _, grade := range []int{1000, 100, 10, 1} {
		n, num = num/grade, num%grade

		// getting each letter for each grade.
		highest = numerals[10*grade]
		lowest = numerals[grade]
		middle = numerals[5*grade]

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

	return buffer.String(), nil
}
