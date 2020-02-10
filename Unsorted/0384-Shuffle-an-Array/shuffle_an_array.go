package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

type Solution struct {
	nums []int // actual input array
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (s *Solution) Reset() []int {
	return s.nums
}

/** Returns a random shuffling of the array. */
func (s *Solution) Shuffle() []int {
	var shuffled = make([]int, len(s.nums))
	copy(shuffled, s.nums)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s.nums), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}

func (s *Solution) String() string {
	return fmt.Sprintf("%v", s.Shuffle())
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func main() {
	var slc = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var sol = Constructor(slc)

	fmt.Printf("Solution: %v\n", sol.String())
	fmt.Printf("Slice:    %v\n", slc)
}
