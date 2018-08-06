package tempconv

import (
	"testing"
)

func Test_Kalvin_Fahrenheit(t *testing.T) {

	temps := map[Kalvin]Fahrenheit{
		Kalvin(255.37): Fahrenheit(0),
		Kalvin(0):      Fahrenheit(-459.67),
		Kalvin(227.59): Fahrenheit(-50.01),
	}

	errorMessage := "Conversation Error %s to %s (Result of KtoF is %s)"

	for k, f := range temps {
		result := KtoF(k)
		if result != f {
			t.Errorf(errorMessage, k, f, result)
		}
	}
}

func Test_Kalvin_Celsius(t *testing.T) {

	temps := map[Kalvin]Celsius{
		Kalvin(273.15): Celsius(0),
		Kalvin(283.15): Celsius(10),
		Kalvin(0):      Celsius(-273.15),
	}

	errorMessage := "Conversation Error %s to %s (Result of KtoC is %s)"

	for k, c := range temps {
		result := KtoC(k)
		if result != c {
			t.Errorf(errorMessage, k, c, result)
		}
	}
}
