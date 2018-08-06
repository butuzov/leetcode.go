package tempconv

import "fmt"

// Celsius temparature
type Celsius float64

const (
	// AbsoluteZeroC is absolute zero temperature in Celcius
	AbsoluteZeroC Celsius = -273.15

	// FreezingC - water freezing point
	FreezingC Celsius = 0

	// BoilingC - water boiling point
	BoilingC Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// CtoF - converts Celsius to Fahrenheit
func CtoF(c Celsius) Fahrenheit {
	if c < AbsoluteZeroC {
		panic("Your input cannot be below absolute zero.")
	}

	//   Conversation formula
	//   t°C * 9/5 + 32
	return Fahrenheit(round(float64(c)*9/5+32, 100))
}

// CtoK - converts Celsius to Kalvin
func CtoK(c Celsius) Kalvin {
	if c < AbsoluteZeroC {
		panic("Your input cannot be below absolute zero.")
	}

	//   Conversation formula
	//   t°C + 273.15 (°K kalvin temperature shift)
	return Kalvin(round(float64(c)+float64(FreezingK), 100))
}
