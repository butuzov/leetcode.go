package prime

import (
	"fmt"
)

func main() {
	fmt.Println(Factors(5001))
}

// Factors return a slice of int64 factors of the provided int`eger.
func Factors(num int64) []int64 {
	factors := []int64{}
	var n int64
	n = 2

	for {

		// do this process while num isnt 0
		if num == 1 {
			break
		}

		// We found one more factor... let's add it to factors list.
		if num%n == 0 || n == num {
			factors = append(factors, n)
			num = int64(num / n)
			continue
		}

		// divisor incrementation

		// NOTE: there should be a nice algorythm to jump over useless
		// divisors. like if number isn't divided by 2, we should skip
		// any number divivded by two.
		// I can't come up with any idea atm. Iterate over existing data?
		// Jump to next prime? How to know whats is next prime?
		n++

	}

	return factors
}
