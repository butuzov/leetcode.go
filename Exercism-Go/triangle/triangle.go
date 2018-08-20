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

// KindFromSides return Kind constant
func KindFromSides(a, b, c float64) Kind {
	// var k Kind

	// At least one of sides less then zero.
	// it's NotATriangle
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	// At least one of sides isNotANumber
	// it's NotATriangle
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}

	// At least one of sides is Infinity
	// it's NotATriangle
	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return NaT
	}

	// All Sides are same.
	// Can we have a problems with floating values? Perhaps.
	avgSide := (a + b + c) / 3
	if avgSide == a && avgSide == b && avgSide == c {
		return Equ
	}

	// Getting Side with Maximum Length + 2 Extra sides.
	// MaxMin will return one (of the if there is two) side with
	// maximum length.
	sides := []float64{a, b, c}
	max, twoSides := MaxMin(sides)

	// Max side is bigger then summ of the other two.
	// This isn't triangle.
	if max > twoSides[0]+twoSides[1] {
		return NaT
	}

	// If two sides are same this is Iso
	for i := 0; i < 3; i++ {
		if sides[i%3] == sides[(i+1)%3] {
			return Iso
		}
	}

	// Fallback - it's jsut triangle.
	return Sca
}

// MaxMin - Return Max lengthed side + slice of other two.
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
