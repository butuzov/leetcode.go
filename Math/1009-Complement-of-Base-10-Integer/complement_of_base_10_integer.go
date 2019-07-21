package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}

	var bits uint
	// counting bits in N by shiffting them to right until we get 0
	for i := N; i > 0; {
		i = i >> 1
		bits++
	}

	// N Bitwise OR  (2^bits)-1 (e.g. 0x1111)
	return N ^ (1<<bits - 1)

}
