package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(i int) (int, error) {
	var step int

	if i < 1 {
		return step, errors.New("Negative number")
	}

	for {
		if i == 1 {
			break
		} else if i%2 == 0 {
			i /= 2
		} else {
			i = i*3 + 1
		}
		step += 1
	}

	return step, nil
}
