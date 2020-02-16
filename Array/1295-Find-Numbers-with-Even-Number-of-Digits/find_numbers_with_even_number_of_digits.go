package main

import (
	"fmt"
	"sync"
)

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func findNumbers(nums []int) int {
	var total int
	for i := 0; i < len(nums); i++ {

		if isOddNumber(nums[i]) {
			total++
		}

	}
	return total
}

func findNumbersWG(nums []int) int {

	type Total struct {
		Count int
		sync.Mutex
	}

	var total Total

	var wg = &sync.WaitGroup{}
	wg.Add(len(nums))

	for i := 0; i < len(nums); i++ {

		go func(n int) {
			defer wg.Done()
			if isOddNumber(n) {
				total.Lock()
				total.Count++
				total.Unlock()
			}
		}(nums[i])
	}

	wg.Wait()

	return total.Count
}

func isOddNumber(n int) bool {
	var counter int
	for n != 0 {
		n = (n - (n % 10)) / 10
		counter++
	}
	return counter%2 == 0
}

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896}))
}
