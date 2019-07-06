package tempconv

import "fmt"

// Fahrenheit temparature.
type Fahrenheit float64

const (
	// AbsoluteZeroF is absolute zero temperature in Fahrenheits
	AbsoluteZeroF Fahrenheit = -459.67

	// FreezingF - freezing temperature
	FreezingF Fahrenheit = 32

	// BoilingF  - water boiling temperature
	BoilingF Fahrenheit = 212
)

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

// FtoC - converts Fahrenheit to Celcius
func FtoC(f Fahrenheit) Celsius {
	if f < AbsoluteZeroF {
		panic("Your input cannot be below absolute zero.")
	}

	//   Conversation formula
	//   (t°F - 32) / 9/5
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
