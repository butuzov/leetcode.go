package main

import (
	"bytes"
	"strings"
)

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var rows = make([]bytes.Buffer, len(s))
	if len(s) > numRows {
		rows = rows[:numRows]
	}

	var b = []byte(s)
	var row int
	var direction bool

	for i := range b {
		rows[row].WriteByte(b[i])
		if row == 0 || row == len(rows)-1 {
			direction = !direction
		}

		if direction {
			row++
		} else {
			row--
		}
	}

	var result string
	for i := range rows {
		result += rows[i].String()
	}

	return result
}

// ---- Maped Bytes

func convertMap(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var rows = make([]bytes.Buffer, len(s))
	if len(s) > numRows {
		rows = rows[:numRows]
	}

	var b = []byte(s)
	var row int
	var direction bool

	for i := range b {
		rows[row].WriteByte(b[i])
		if row == 0 || row == len(rows)-1 {
			direction = !direction
		}

		if direction {
			row++
		} else {
			row--
		}
	}

	return strings.Join(func(b []bytes.Buffer) []string {
		var s = make([]string, len(b))
		for i := range b {
			s[i] = b[i].String()
		}
		return s
	}(rows), "")
}
