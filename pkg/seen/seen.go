/*
	Package seen is for seen lookups
*/

package seen

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Seen struct {
	rows int
	cols int

	seen []bool
}

func (s Seen) IsSeen(row, col int) bool {
	return s.seen[s.Index(row, col)]
}

func (s Seen) String() string {
	var buf bytes.Buffer

	// index cell
	leftColWidth := len(strconv.Itoa(s.rows))
	fmt.Println(leftColWidth)
	// header
	buf.WriteString(strings.Repeat(" ", leftColWidth))
	for i := 0; i < s.cols; i++ {
		fmt.Fprintf(&buf, "|%d", i)
	}
	LineSize := buf.Len()
	fmt.Fprintln(&buf)
	fmt.Fprintln(&buf, strings.Repeat("-", LineSize))
	for i := 0; i < s.rows; i++ {
		fmt.Fprintf(&buf, fmt.Sprintf("%%-%dd", leftColWidth), i)

		for j := 0; j < s.cols; j++ {

			// TODO: cach this call.
			width := len(strconv.Itoa(j))
			if s.IsSeen(i, j) {
				fmt.Fprintf(&buf, fmt.Sprintf("|%%-%ds", width), "x")
			} else {
				fmt.Fprintf(&buf, fmt.Sprintf("|%%-%ds", width), " ")
			}
		}
		fmt.Fprintln(&buf)
	}

	return buf.String()
}

func (s Seen) Seen(row, col int) {
	s.seen[s.Index(row, col)] = true
}

func (s Seen) Index(row, col int) int {
	idx := row*s.cols + col
	return idx
}

func Look(row, col int) (*Seen, error) {
	if row <= 0 || col <= 0 {
		return nil, fmt.Errorf("Row(%d)/Col(%d) below or eq zero", row, col)
	}
	return &Seen{
		rows: row,
		cols: col,
		seen: make([]bool, row*col),
	}, nil
}
