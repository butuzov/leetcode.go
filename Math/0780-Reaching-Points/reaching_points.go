package main
/*******************************************************************************
  Problem Solution
*******************************************************************************/

func reachingPoints(sx int, sy int, tx int, ty int) bool {

	if sx > tx || sy > ty {
		return false
	}

	// we slowly (becouse of decrementation) moving down to sx&sy
	// its slow but easy to understand. You can speed up this by adding
	// modulus operator.
	for (tx > sx) || (ty > sy) {
		if tx == ty {
			break
		}
		if tx > ty {
			tx -= ty
			continue
		}

		if ty > tx {
			ty -= tx
			continue
		}
	}

	return sx == tx && sy == ty
}
