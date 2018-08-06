package tempconv

import "math"

// general helper function - round numbers to next .00 point.
func round(n, unit float64) float64 {
	return math.Round(n*unit) / unit
}
