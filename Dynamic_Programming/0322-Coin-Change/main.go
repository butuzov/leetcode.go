package main

import (
	"fmt"
)

func main() {
	// dude, write some code...
	fmt.Printf("%d\n", coinChange([]int{1}, 1))
}

func coinChange(coins []int, amount int) int {
	max := 10000

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = max
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}

			// dp
			if dp[i-coin]+1 <= dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	if dp[len(dp)-1] == max {
		return -1
	}

	return dp[len(dp)-1]
}
