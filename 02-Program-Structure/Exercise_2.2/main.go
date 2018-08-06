package main

import (
	"fmt"

	tc "./tempconv"
)

func main() {
	fmt.Println("This code converts temperature in Celcius/Fahrenheiht/Kalvin (e.g.)")
	fmt.Printf("%5s -> %9s \n", tc.Celsius(0), tc.CtoF(tc.Celsius(10)))
	fmt.Printf("%5s -> %9s \n", tc.Celsius(10), tc.CtoK(tc.Celsius(10)))

	fmt.Printf("%5s -> %9s \n", tc.Kalvin(10), tc.KtoC(tc.Kalvin(10)))
	fmt.Printf("%5s -> %9s \n", tc.Kalvin(0), tc.KtoF(tc.Kalvin(0)))

	fmt.Printf("%5s -> %9s \n", tc.Fahrenheit(10), tc.FtoC(tc.Fahrenheit(10)))
	fmt.Printf("%5s -> %9s \n", tc.Fahrenheit(0), tc.FtoK(tc.Fahrenheit(0)))
}
