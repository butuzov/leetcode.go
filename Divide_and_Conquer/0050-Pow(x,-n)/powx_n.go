package main
/*******************************************************************************
  Problem Solution
*******************************************************************************/

// magic numbers
const (
	ANY_POWER_ZERO = 1
	ONE_POWER_ANY  = 1
	POWER_ZERO     = 0
	POWER_ONE      = 1
)

// Moemoize Key Type
type params struct {
	x float64
	n int
}

// Memoize accpets callback and cache results based on arguments.
// but we need to
func Memoize(f func(x float64, n int) float64) func(x float64, n int) float64 {

	var cache = make(map[params]float64)

	return func(x float64, n int) float64 {

		var Param = params{x, n}

		if _, ok := cache[Param]; !ok {
			cache[Param] = f(x, n)
		}

		return cache[Param]
	}

}

func myPow(x float64, n int) float64 {

	// Predefining Power function
	var power func(x float64, n int) float64

	// Momoizing POwer
	power = Memoize(func(x float64, n int) float64 {

		// self call.
		if n == POWER_ONE {
			return x
		}

		// n is 0
		if n == POWER_ZERO {
			return ANY_POWER_ZERO
		}

		// x is 1
		if x == ONE_POWER_ANY {
			return ONE_POWER_ANY
		}

		// if x and n (lets say) 3^5 (3 and 5 respectevly)
		// we form a sentanse of 3 * 3^(int(5/2)) * 3^(int(5/2))
		// which is 3 * 3^2 * 3^2
		var multiplier float64 = 1.0

		if n%2 == 1 {
			multiplier = x
		}

		// default downsize case
		return multiplier * power(x, n/2) * power(x, n/2)
	})

	// regular powering
	if n < 0 {
		return 1 / power(x, -n)
	}
	return power(x, n)
}
