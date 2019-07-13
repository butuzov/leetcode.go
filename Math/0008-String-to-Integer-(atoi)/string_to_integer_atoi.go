package main

import (
	"errors"
	"math"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func myAtoi(str string) int {

	// Converts string to bytes slice
	var strToBytes = func(s string, res *[]byte) (int, error) {

		var isPlus, isSpace, isMinus, isNum bool

		var isNumerInSequence = false

		for i := 0; i < len(s); i++ {
			isPlus = s[i] == '+'
			isSpace = s[i] == ' '
			isMinus = s[i] == '-'
			isNum = (s[i] >= '0' && s[i] <= '9')

			// yeah, it's a space
			if len(*res) == 0 && isSpace {
				continue
			}

			// begining of valid sequence
			if len(*res) == 0 && (isPlus || isMinus || isNum) {
				isNumerInSequence = isNum
				*res = append(*res, s[i])
				continue
			}

			// not valid
			if len(*res) == 0 && !(isPlus || isMinus || isNum) {
				return 0, errors.New("No valid sequense on start")
			}

			if len(*res) != 0 && isNum {
				isNumerInSequence = isNum
				*res = append(*res, s[i])
				continue
			}

			if len(*res) != 0 && !isNum && !isNumerInSequence {
				return 0, errors.New("No valid sequense on start")
			}

			// and here we will stop.
			if len(*res) != 0 && !isNum && isNumerInSequence {
				break
			}
		}

		if len(*res) == 1 && !isNumerInSequence {
			return 1, errors.New("No valid sequense on start")
		}

		return len(*res), nil

	}

	// overflown return status of the number
	// by throwing OK status, if number is actually overflown.
	// input result * mul
	// for the next step please do input /= nul
	var overflown = func(n int64) bool {
		if n > math.MaxInt32 {
			return true
		}

		if n < math.MinInt32 {
			return true
		}
		return false
	}

	var storage []byte
	var n, error = strToBytes(str, &storage)
	if n == 0 || error != nil {
		return 0
	}

	var result int64

	// getting sign from the bytes array.
	var multiplier = int64(1)
	if storage[0] == '-' {
		multiplier = int64(-1)
	}
	if storage[0] == '-' || storage[0] == '+' {
		storage = storage[1:]
	}
	if len(storage) == 0 {
		return 0
	}

	// no zeros on start
	for storage[0] == '0' {
		if len(storage) == 1 {
			storage = []byte{}
			break
		}
		storage = storage[1:]
	}

	if len(storage) == 0 {
		return 0
	}

	var pow int64
	pow = 1

	// fmt.Println(storage)
	for i := len(storage) - 1; i >= 0; i-- {

		if overflown(result + int64(storage[i]-'0')*pow) {
			switch multiplier {
			case -1:
				return math.MinInt32
			case 1:
				return math.MaxInt32
			}
		}

		result += int64(storage[i]-'0') * pow

		if i > 0 {
			if overflown(pow * 10) {
				switch multiplier {
				case -1:
					return math.MinInt32
				case 1:
					return math.MaxInt32
				}
			}
		}
		pow *= 10
	}

	return int(result * multiplier)
}
