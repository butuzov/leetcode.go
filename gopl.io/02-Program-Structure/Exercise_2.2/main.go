package main

import (
	"flag"
	"fmt"
	"log"

	tc "./tempconv"
)

// Cumulative arguments.

// Flags arguments.
var temperature = flag.Float64("t", 0, "A Feed Token from TeemTreeHouse.com")
var scale = flag.String("s", "Celsius", "Temperature Scale can be Celsius, Fahrenheit or Kalvin")

func main() {
	flag.Parse()

	var errorBelowZeroMsg = "Provided temperature %s below absolute zero point\n"
	var reportMsg = "%s is %s or %s\n"

	fmt.Printf("You have entered: %0.2f °%s\n", *temperature, *scale)

	switch *scale {
	case "Celsius":

		degrees := tc.Celsius(*temperature)

		if degrees < tc.AbsoluteZeroC {
			log.Fatalf(errorBelowZeroMsg, degrees)
		}

		fmt.Printf(reportMsg, degrees, tc.CtoF(degrees), tc.CtoK(degrees))

		break
	case "Fahrenheit":

		degrees := tc.Fahrenheit(*temperature)

		if degrees < tc.AbsoluteZeroF {
			log.Fatalf(errorBelowZeroMsg, degrees)
		}

		fmt.Printf(reportMsg, degrees, tc.FtoC(degrees), tc.FtoK(degrees))

		break
	case "Kalvin":

		degrees := tc.Kalvin(*temperature)

		if degrees < tc.AbsoluteZeroK {
			log.Fatalf(errorBelowZeroMsg, degrees)
		}

		fmt.Printf(reportMsg, degrees, tc.KtoF(degrees), tc.KtoC(degrees))

		break

	}

	// fmt.Println("This code converts temperature in Celcius/Fahrenheiht/Kalvin (e.g.)")
	// fmt.Printf("%5s -> %9s \n", tc.Celsius(0), tc.CtoF(tc.Celsius(10)))
	// fmt.Printf("%5s -> %9s \n", tc.Celsius(10), tc.CtoK(tc.Celsius(10)))

	// fmt.Printf("%5s -> %9s \n", tc.Kalvin(10), tc.KtoC(tc.Kalvin(10)))
	// fmt.Printf("%5s -> %9s \n", tc.Kalvin(0), tc.KtoF(tc.Kalvin(0)))

	// fmt.Printf("%5s -> %9s \n", tc.Fahrenheit(10), tc.FtoC(tc.Fahrenheit(10)))
	// fmt.Printf("%5s -> %9s \n", tc.Fahrenheit(0), tc.FtoK(tc.Fahrenheit(0)))
}
