package hamming

import "errors"

// Distance is simple calucator of Hamming distance betwean DNA
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("DNA Sequence isn't same length")
	}

	var distance int

	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
