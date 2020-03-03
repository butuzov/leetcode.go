package main

import "fmt"

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func maxProfit(prices []int) (max int) {

	if prices == nil || len(prices) <= 1 {
		return 0
	}

	// var max int         // (max) profit
	// var min = prices[0] // min price

	for i := range prices {
		if prices[i] < prices[0] {
			prices[0] = prices[i]
		} else if prices[i]-prices[0] > max {
			max = prices[i] - prices[0]
		}
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 10}))
}
