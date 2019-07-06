package tempconv

import (
	"testing"
)

func Test_Fahrenheit_Celcius(t *testing.T) {

	temps := map[Fahrenheit]Celsius{
		Fahrenheit(50):      Celsius(10),
		Fahrenheit(32):      Celsius(0),
		Fahrenheit(-148):    Celsius(-100),
		Fahrenheit(-459.67): Celsius(-273.15),
	}

	errorMessage := "Conversation Error %s to %s (Result of FtoC is %s)"

	for f, c := range temps {
		result := FtoC(f)
		if result != c {
			t.Errorf(errorMessage, f, c, result)
		}
	}
}

func Test_Fahrenheit_Kalvin(t *testing.T) {

	temps := map[Fahrenheit]Kalvin{
		Fahrenheit(0):       Kalvin(255.37),
		Fahrenheit(-459.67): Kalvin(0),
		Fahrenheit(-50):     Kalvin(227.59),
	}

	errorMessage := "Conversation Error %s to %s (Result of FtoK is %s)"

	for f, k := range temps {
		result := FtoK(f)
		if result != k {
			t.Errorf(errorMessage, f, k, result)
		}
	}
}
