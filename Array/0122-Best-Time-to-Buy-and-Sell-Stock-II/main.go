package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func maxProfit(prices []int) int {
	var profit int
	if len(prices) <= 1 {
		return profit
	}

	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}
