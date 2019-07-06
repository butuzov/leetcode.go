package collatzconjecture

import (
	"errors"
)

// CollatzConjecture return number of steps required
// to solve Collatz Conjecture problem.
func CollatzConjecture(i int) (int, error) {
	var step int

	// exit condition "one"...
	if i < 1 {
		return step, errors.New("Negative number")
	}

	for {
		if i == 1 {
			// exit condition
			break
		} else if i%2 == 0 {
			// Requirment 1
			i /= 2
		} else {
			// Requirment 2
			i = i*3 + 1
		}
		step++
	}

	return step, nil
}
