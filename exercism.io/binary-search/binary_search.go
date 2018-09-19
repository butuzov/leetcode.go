package binarysearch

import (
	"fmt"
)

// SearchInts os binanry search implementation.
// returns -1 if match not found.
func SearchInts(slice []int, key int) int {

	if len(slice) == 0 {
		return -1
	}

	// Boudaries and Index
	var low, up, idx int
	low, up = 0, len(slice)-1

	for {

		// Fail state.
		if low > up {
			return -1
		}

		// Index to check
		idx = (up + low) / 2

		switch {
		case slice[idx] == key:
			// Match found!
			return idx
		case slice[idx] > key:
			// Moving up boundary to current index
			up = idx - 1
			continue
		case slice[idx] < key:
			// Moving low boundary to current index
			low = idx + 1
			continue
		}
	}

	// this is unreachable code, but logick require me to put it here.
	return -1
}

func main() {
	for _, tC := range testCases {
		fmt.Printf("%#v\n", tC.slice)
		found := SearchInts(tC.slice, tC.key)
		fmt.Printf("Search %5d / Index  %3d == %3d\n\n", tC.key, tC.x, found)
	}
}
