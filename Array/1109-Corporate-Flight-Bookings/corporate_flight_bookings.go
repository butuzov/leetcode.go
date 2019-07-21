package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func corpFlightBookings(bookings [][]int, n int) []int {
	var results = make([]int, n+1)
	for _, v := range bookings {
		results[v[0]-1] += v[2] // in
		results[v[1]] -= v[2]   // out
	}

	for i := 0; i < n-1; i++ {
		results[i+1] += results[i]
	}

	return results[:len(results)-1]
}
