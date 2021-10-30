package main

import (
	"fmt"
	"time"
)

func main() {
	// dude, write some code...
	fmt.Println("data", dayOfYear("2003-03-01"))
}

func dayOfYear(date string) int {
	t, _ := time.Parse("2006-01-02", date)
	diff := t.Sub(time.Date(t.Year(), 1, 0, 0, 0, 0, 0, time.UTC))
	return int(diff / (24 * time.Hour))
}
