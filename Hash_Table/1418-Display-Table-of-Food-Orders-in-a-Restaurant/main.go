package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {

	fmt.Println(TestCases[0].orders)
	fmt.Println(strings.Repeat("===", 20))
	fmt.Println(TestCases[0].expected)
	fmt.Println(strings.Repeat("===", 20))
	fmt.Println(displayTable(TestCases[0].orders))
	log.Printf("2222 %s", "aaa")

}
