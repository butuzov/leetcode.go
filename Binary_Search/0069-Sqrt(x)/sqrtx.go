package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/

// https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Babylonian_method
func Babylonian(n int) int {

	var (
		o = float64(n) // Original value as float64
		x = float64(n) // x of binary search
		y = 1.0        // y of binary search
		e = 1e-5       // error
	)

	for x-y > e {
		x = (x + y) / 2
		y = o / x
		// fmt.Printf("y=%f, x=%f\n", y, x)
	}

	return int(x)
}

// Bakhshali - https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Bakhshali_method
func Bakhshali(n int) int {

	var iterate = func(x float64) float64 {
		a := (float64(n) - x*x) / (2 * x)
		xa := x + a
		return xa - ((a * a) / (2 * xa))
	}

	var (
		o = float64(n)
		x = float64(n) / 2.0
		e = 1e-5
	)

	for x*x-o > e {
		x = iterate(x)
	}
	return int(x)
}

// Fastest method

func mySqrt(n int) int {
	return Bits(n)
}

func Bits(x int) int {

	var res int
	var bit = 1 << 30

	for bit > x {
		bit >>= 2
	}

	for bit != 0 {
		if x >= res+bit {
			x -= res + bit
			res = (res >> 1) + bit
		} else {
			res >>= 1
		}
		bit >>= 2
	}

	return res
}
