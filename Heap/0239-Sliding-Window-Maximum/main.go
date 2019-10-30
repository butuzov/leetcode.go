package main

import (
	"fmt"
	"strings"
)

func debug(in []int) {
	fmt.Println(strings.Repeat("=", 100))

	for i, _ := range in {
		fmt.Printf("%4d ", i)
	}
	fmt.Println()

	fmt.Println(strings.Repeat("=", 100))

	for _, i := range in {
		fmt.Printf("%4d ", i)
	}
	fmt.Println()

	fmt.Println(strings.Repeat("=", 100))
}

func main() {

	var in = []int{-6, -10, -7, -1, -9, 9, -8, -4, 10, -5, 2, 9, 0, -7, 7, 4, -2, -10, 8, 7}
	var k = 7

	fmt.Printf("k - %d, Incoming slice\n", k)
	debug(in)

	fmt.Println("Results   ", maxSlidingWindow(in, k))

	fmt.Println("Exapected ", []int{9, 9, 10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 8, 8})
}
