package tempconv

import "fmt"

// Kalvin temparature
type Kalvin float64

const (
	// AbsoluteZeroK is absolute zero temperature in Kalvins
	AbsoluteZeroK Kalvin = 0
	// FreezingK - freezing temperature
	FreezingK Kalvin = 273.15
	// BoilingK  - water boiling temperature
	BoilingK Kalvin = 373.15
)

func (k Kalvin) String() string {
	return fmt.Sprintf("%g°K", k)
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
