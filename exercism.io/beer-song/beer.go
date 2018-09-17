package beer

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode"
)

func main() {
	verses, err := Verses(20, 0)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(verses)
}

// Conjugate - number to string (buttles number)
func Conjugate(bottles int) string {

	switch {
	case bottles == 0:
		return "no more bottles"
	case bottles == 1:
		return "1 bottle"
	}

	return fmt.Sprintf("%d bottles", bottles)
}

// Song generate full 99 botles song.
func Song() string {
	song, err := Verses(99, 0)
	if err != nil {
		log.Fatalln(err)
	}
	return song
}

// Verses - generate part of song based on low and up number.
func Verses(up, low int) (string, error) {
	if up < low {
		return "", errors.New("Start is much larger than stop")
	}

	// Failing if arguments can't satisfy conditions
	if (up >= 100 || up < 0) || (low >= 100 || low < 0) {
		return "", errors.New("calling Verses(99,0)")
	}

	// New bytes buffer
	verses := bytes.NewBuffer([]byte{})

	// Generating Verses
	for i := up; i >= low; i-- {
		verse, err := Verse(i)
		if err != nil {
			log.Fatal(err)
		}
		verses.WriteString(verse)
		verses.WriteString("\n")
	}

	return verses.String(), nil
}

// Verse generates a verse based on number.
func Verse(verses int) (string, error) {
	if verses >= 100 || verses < 0 {
		return "", errors.New("calling Verses(99,0)")
	}

	// Getting this and previous numerals as strings (with buttles! string)
	bottles := Conjugate(verses)
	left := Conjugate(verses - 1)

	take := "one"
	if verses == 1 {
		take = "it"
	}

	var lines []string

	// We changing first parameter to Uppercase,
	// we need this only for 'no more buttles' left.
	lines = append(lines, fmt.Sprintf("%s of beer on the wall, %s of beer.\n", Ucfirst(bottles), bottles))

	// Second line - number of buttles left.
	if verses == 0 {
		lines = append(lines, "Go to the store and buy some more, 99 bottles of beer on the wall.\n")
	} else {
		lines = append(lines, fmt.Sprintf("Take %s down and pass it around, %s of beer on the wall.\n", take, left))
	}

	return strings.Join(lines[:], ""), nil
}

// This is code function I got from
// https://www.php2golang.com/method/function.ucfirst.html
// maybe it's better just to check is first byte is 'n'
func Ucfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}
