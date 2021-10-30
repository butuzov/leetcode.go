package main

import (
	"fmt"
	"time"
)

func main() {
	// dude, write some code...
	fmt.Println(daysBetweenDates("2019-12-31", "2020-01-15"))
}

func daysBetweenDates(date1 string, date2 string) int {
	t1, _ := time.Parse("2006-01-02", date1)
	t2, _ := time.Parse("2006-01-02", date2)

	diff := t1.Sub(t2)
	if diff < 0 {
		diff *= -1
	}

	return int(diff / (time.Hour * 24))
}
