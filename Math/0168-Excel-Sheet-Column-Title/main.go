package main

import (
	"fmt"
)

func main() {
	for _, v := range TestCases {
		fmt.Println(v.expected, v.n, convertToTitle(v.n))
		fmt.Println("")
	}
}
