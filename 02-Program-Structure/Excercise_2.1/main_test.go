package main

import (
	"testing"
)

func Test_Celcius_Fahrenheit(t *testing.T) {

	temps := map[Celsius]Fahrenheit{
		Celsius(10):   Fahrenheit(50),
		Celsius(0):    Fahrenheit(32),
		Celsius(-100): Fahrenheit(-148),
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
