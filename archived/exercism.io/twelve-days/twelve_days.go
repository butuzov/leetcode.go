package twelve

import (
	"bytes"
	"fmt"
	"strings"
)

// Song return imploded by "\n" verses.
func Song() string {
	b := bytes.NewBuffer([]byte{})
	for i := 1; i <= 12; i++ {
		b.WriteString(Verse(i) + "\n")
	}
	return b.String()
}

// Verse Generated for the song.
func Verse(n int) string {

	if n > 12 {
		return ""
	}

	var verses = []struct {
		Day     string
		Present string
	}{
		{"first", "a Partridge in a Pear Tree"},
		{"second", "two Turtle Doves"},
		{"third", "three French Hens"},
		{"fourth", "four Calling Birds"},
		{"fifth", "five Gold Rings"},
		{"sixth", "six Geese-a-Laying"},
		{"seventh", "seven Swans-a-Swimming"},
		{"eighth", "eight Maids-a-Milking"},
		{"ninth", "nine Ladies Dancing"},
		{"tenth", "ten Lords-a-Leaping"},
		{"eleventh", "eleven Pipers Piping"},
		{"twelfth", "twelve Drummers Drumming"},
	}

	var out []string
	for i := (n - 1); i >= 0; i-- {
		if n > 1 && i == 0 {
			out = append(out, "and "+verses[i].Present)
			continue
		}
		out = append(out, verses[i].Present)
	}

	message := "On the %s day of Christmas my true love gave to me, %s."
	return fmt.Sprintf(message, verses[n-1].Day, strings.Join(out, ", "))
}
