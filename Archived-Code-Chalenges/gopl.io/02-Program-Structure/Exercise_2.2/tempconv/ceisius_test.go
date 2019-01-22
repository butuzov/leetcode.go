package tempconv

import (
	"testing"
)

func Test_Celcius_Fahrenheit(t *testing.T) {

	temps := map[Celsius]Fahrenheit{
		Celsius(10):   Fahrenheit(50),
		Celsius(0):    Fahrenheit(32),
		Celsius(-100): Fahrenheit(-148),
		Celsius(100):  Fahrenheit(212),
	}

	errorMessage := "Conversation Error %s to %s (Result of CtoF is %s)"

	for c, f := range temps {
		result := CtoF(c)
		if result != f {
			t.Errorf(errorMessage, c, f, result)
		}
	}
}

func Test_Celcius_Kalvin(t *testing.T) {

	temps := map[Celsius]Kalvin{
		Celsius(-273.15): Kalvin(0),
		Celsius(0):       Kalvin(273.15),
		Celsius(-100):    Kalvin(173.15),
	}

	errorMessage := "Conversation Error %s to %s (Result of CtoK is %s)"

	for c, k := range temps {
		result := CtoK(c)
		if result != k {
			t.Errorf(errorMessage, c, k, result)
		}
	}
}
