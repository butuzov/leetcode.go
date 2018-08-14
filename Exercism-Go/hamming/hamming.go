package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("DNA Sequence isn't same length")
	}

	var distance int

	for i := range a {
		if a[i] != b[i] {
			distance += 1
		}
	}

	return distance, nil
}
