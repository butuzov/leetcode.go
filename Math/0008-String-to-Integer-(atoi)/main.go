package main

import (
	"fmt"
)

func main() {

	i := myAtoi("2147483646")
	fmt.Println(i == 2147483646)
}
