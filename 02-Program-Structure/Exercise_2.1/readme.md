# Exercise 2.1

Add types, constants and functions to `tempconv` for processing temperatures in Kelvin scale, where zero Kelvin is -273.15 °C and differenc of 1K the same magnitude as 1 °C.


```go
	fmt.Printf("%5s -> %9s \n", Celsius(10), CtoF(Celsius(10)))
	fmt.Printf("%5s -> %9s \n", Celsius(10), CtoK(Celsius(10)))

	fmt.Printf("%5s -> %9s \n", Kalvin(10), KtoC(Kalvin(10)))
	fmt.Printf("%5s -> %9s \n", Kalvin(0), KtoF(Kalvin(0)))

	fmt.Printf("%5s -> %9s \n", Fahrenheit(10), FtoC(Fahrenheit(10)))
	fmt.Printf("%5s -> %9s \n", Fahrenheit(0), FtoK(Fahrenheit(0)))
```
