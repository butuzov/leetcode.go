package prime

// Nth search for n-th prime number
func Nth(n int) (int, bool) {

	if n < 1 {
		return 0, false
	}

	var primes []int
	var number = 1
	var isSolution bool
	var isPrime = false
	for {
		number++

		isPrime = true

		for i := range primes {
			div := number % primes[i]
			if div == 0 {
				isPrime = false
				break
			}
		}

		if isPrime == false {
			continue
		} else {
			primes = append(primes, number)

			if len(primes) == n {
				isSolution = true
			}
		}

		if isSolution {
			break
		}
	}
	return primes[n-1], true
}
