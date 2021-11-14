package main

import "fmt"

func main() {
	// dude, write some code...
	fmt.Printf("%d", change(5, []int{1, 2, 5}))
}

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}
	return dp[amount]
}
