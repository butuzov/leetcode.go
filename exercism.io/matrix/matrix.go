package matrix

import (
	"bytes"
	"errors"
	"math"
	"strconv"
)

// Matrix struct
type Matrix struct {
	rows     int    // Rows Number
	cols     int    // Cols Number
	values   []int  // Onedimential Matrix representation
	original string // Original String
}

// New creates new matrix based on string representation.
func New(s string) (*Matrix, error) {
	matrix := new(Matrix)
	matrix.original = s

	// Position indicators
	var row, col int
	var isRow, isCol bool
	// var skipSpaceCheck bool

	// buffer
	var b bytes.Buffer

	// we assuming that string is ascii characters only.
	for i := 0; i < len(s); i++ {

		// we matched some OK bytes
		if (s[i] >= 48 && s[i] <= 57) || (s[i] == 32 || s[i] == 10) {

			// We have a case of starting space(s) after \n byte, we simply
			// ignoring it.
			if col == 0 && s[i] == 32 {
				continue
			}

			// This tracks the the row and column.
			if isRow == false {
				isRow = true
				row++
			}

			// Tracks new column
			if isCol == false {
				isCol = true
				col++
			}

			// finally - digit!
			if s[i] >= 48 && s[i] <= 57 {
				b.WriteByte(s[i])
			}

			// if this a space or a new line or its last element of string.
			if s[i] == 32 || s[i] == 10 || i == (len(s)-1) {

				// convert bufferToNumber
				d, err := bufferToNumber(&b)

				if err != nil {
					return matrix, errors.New("This number isn't ok")
				}

				// reset
				b.Reset()

				// adding number value.
				matrix.values = append(matrix.values, d)

				// after digit added, we marking that we need to check a new column
				isCol = false

				// After we found a new line or its a end
				if s[i] == 10 || i == (len(s)-1) {

					matrix.rows++

					if matrix.cols == 0 {
						matrix.cols = col
					} else if col != matrix.cols {
						return matrix, errors.New("different number of cols")
					}

					// inticator for new row
					if s[i] == 10 {
						isRow = false
					}
					col = 0
				}
			}
		} else {
			// not a space, not a new line, not a number
			return matrix, errors.New("empty line")
		}
	}

	// new row was indicated but no new
	if isRow == false {
		return matrix, errors.New("New row declared, but it's empty...")
	}

	return matrix, nil
}

// bufferToNumber converts buffer to string by multiplication
// of 10 each of numbers in bytes.
// its a bit different then jsut run strconv package
func bufferToNumber(b *bytes.Buffer) (int, error) {
	var number int
	bytes, length := b.Bytes(), b.Len()-1

	for i := 0; i <= length; i++ {
		number += int(bytes[i]-48) * int(math.Pow10(length-i))
	}

	if number < 0 {
		return 0, errors.New("ccc")
	}

	return number, nil
}

// String represents matrix as matrix.
// debug only
func (m *Matrix) String() string {
	var b bytes.Buffer
	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.cols; col++ {

			// not going convert int to string using
			// ascii table and powers of number, this is
			// just a simple matrix representation after all.
			b.WriteString(strconv.Itoa(m.values[row*m.cols+col]))

			// space
			if col != (m.cols - 1) {
				b.WriteByte(32)
			}
		}
		// new line
		b.WriteByte(10)
	}
	return b.String()
}

// Rows return slice of row slices
func (m *Matrix) Rows() [][]int {
	var out [][]int

	for row := 0; row < m.rows; row++ {
		out = append(out, make([]int, 0))
		for col := 0; col < m.cols; col++ {
			out[row] = append(out[row], m.values[row*m.cols+col])
		}
	}

	return out
}

// Set return ok, if value set. !ok if there was an error.
func (m *Matrix) Set(row, col, val int) bool {

	if (row < 0 || m.rows <= row) || (col < 0 || m.cols <= col) {
		return false
	}

	if row*m.cols+col > (m.rows*m.cols - 1) {
		return false
	}

	m.values[row*m.cols+col] = val

	// fmt.Printf("Matrix After\n%s\n", m)

	return true
}

// Cols return slice of cols slices - O(n)
func (m *Matrix) Cols() [][]int {
	var out [][]int
	for col := 0; col < m.cols; col++ {
		out = append(out, make([]int, 0))
		for row := 0; row < m.rows; row++ {
			out[col] = append(out[col], m.values[row*m.cols+col])
		}
	}
	return out
}
