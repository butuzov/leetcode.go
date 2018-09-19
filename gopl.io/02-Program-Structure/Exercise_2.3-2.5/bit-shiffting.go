package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b uint64
	a, b = 1311768467294899695, 4

	var mapa map[string]uint64 = map[string]uint64{
		"!var a": a,
		"!var b": b,

		"and": a & b,
		"or":  a | b,
		"xor": a ^ b,

		// each bit of z is 0
		// - if corresponding bit of y is 1,
		// - else corresponding bit of x
		"bit clear": a &^ b,

		// left shift
		"left shift": a << b,

		// right shift
		"right shift": a >> b,
	}

	for k, v := range mapa {
		fmt.Printf("%-16s %.64b\n", k, v)
	}

	fmt.Printf("%d\n", 0x1234567890ABCDEF)

	fmt.Println(strings.Repeat("-", 40))

	// fmt.Printf("%-8s %d %d\n", "testing", a, a&(a-1))
	// fmt.Printf("%-8s %d %d\n", "testing", a, a|(a+1))
	// fmt.Printf("%-8s %d %d\n", "testing", a, a&(a+1))
	// fmt.Printf("%-8s %d %d\n", "testing", a, a&(-a))
	// fmt.Printf("%-16s %.8b\n", "testing", a&(a-1))
	// fmt.Printf("%-16s %.8b\n", "testing", 2)

}
