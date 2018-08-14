// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// var k Kind

	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}

	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return NaT
	}

	// All Sides are same. Can we have a problems with floating values?
	// perhaps.
	avg_side := (a + b + c) / 3
	if avg_side == a && avg_side == b && avg_side == c {
		return Equ
	}

	// Getting Side with Maximum Length + 2 Extra sides.
	sides := []float64{a, b, c}
	max, two_sides := MaxMin(sides)

	// Max is bigger then summ of the other two.
	// this isn't triangle
	if max > two_sides[0]+two_sides[1] {
		return NaT
	}

	// if two sides are same this is Iso
	for i := 0; i < 3; i++ {
		if sides[i%3] == sides[(i+1)%3] {
			return Iso
		}
	}

	// Fallback
	return Sca
}

// Return Max lengthed side + slice of other two.
func MaxMin(s []float64) (float64, []float64) {
	var max float64
	var index int

	for i, value := range s {
		if max < value {
			max, index = value, i
		}
	}
	return max, []float64{s[(index+1)%3], s[(index+2)%3]}
}
