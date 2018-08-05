package main

import (
	"fmt"
	"math"
)

// Celsius temparature
type Celsius float64

// Fahrenheit temparature.
type Fahrenheit float64

// Kalvin temparature
type Kalvin float64

func main() {
	fmt.Println("This code converts temperature in Celcius/Fahrenheiht/Kalvin (e.g.)")
	fmt.Printf("%5s -> %9s \n", Celsius(10), CtoF(Celsius(10)))
	fmt.Printf("%5s -> %9s \n", Celsius(10), CtoK(Celsius(10)))

	fmt.Printf("%5s -> %9s \n", Kalvin(10), KtoC(Kalvin(10)))
	fmt.Printf("%5s -> %9s \n", Kalvin(0), KtoF(Kalvin(0)))

	fmt.Printf("%5s -> %9s \n", Fahrenheit(10), FtoC(Fahrenheit(10)))
	fmt.Printf("%5s -> %9s \n", Fahrenheit(0), FtoK(Fahrenheit(0)))
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100

	AbsoluteZeroK Kalvin = 0
	FreezingK     Kalvin = 273.15
	BoilingK      Kalvin = 373.15

	AbsoluteZeroF Fahrenheit = -459.67
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kalvin) String() string     { return fmt.Sprintf("%g°K", k) }

func round(n, unit float64) float64 {
	return math.Round(n*unit) / unit
}

// CtoF - converts Celsius to Fahrenheit
func CtoF(c Celsius) Fahrenheit {
	if c < AbsoluteZeroC {
		panic("Your input cannot be below absolute zero.")
	}

	return Fahrenheit(round(float64(c)*9/5+32, 100))
}

// CtoK - converts Celsius to Kalvin
func CtoK(c Celsius) Kalvin {
	if c < AbsoluteZeroC {
		panic("Your input cannot be below absolute zero.")
	}

	return Kalvin(round(float64(c)+float64(FreezingK), 100))
}

// FtoC - converts Fahrenheit to Celcius
func FtoC(f Fahrenheit) Celsius {
	if f < AbsoluteZeroF {
		panic("Your input cannot be below absolute zero.")
	}

	return Celsius(round((float64(f)-32)*5/9, 100))
}

// FtoK - converts Fahrenheit to Kalvin
func FtoK(f Fahrenheit) Kalvin {
	if f < AbsoluteZeroF {
		panic("Your input cannot be below absolute zero.")
	}
	//   Conversation formula
	//   (t°F - 32) / 9/5 + 273.15

	return Kalvin(round((float64(f)-32)/1.8+273.15, 100))
}

// KtoC - converts Kalvin to Celcius
func KtoC(k Kalvin) Celsius {
	if k < AbsoluteZeroK {
		panic("Your input cannot be below absolute zero.")
	}
	//   Conversation formula
	//   t°K - 273.15

	return Celsius(round(float64(k)-273.15, 100))
}

// KtoF - converts Kalvin to Fahrenheit
func KtoF(k Kalvin) Fahrenheit {
	if k < AbsoluteZeroK {
		panic("Your input cannot be below absolute zero.")
	}

	//   Conversation formula
	//   (t°K - 273.15) * 9/5 + 32

	return Fahrenheit(round((float64(k)-273.15)*1.8+32, 100))
}
