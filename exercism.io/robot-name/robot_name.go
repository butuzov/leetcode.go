package robotname

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// Ways to solve this problems

// 1: Generate Each time new number. Keep cache of names and check each time
//    if there collisions. If one found - regenerate. Once we cache 26*26*1000
//    names in our cache we doooooooooomed.
// 2: Randomly generate Two letters and use them as key in map where we going
//    to store generated robot numbers. It's way more optimal for look up, but
//    still require map iteration.
// 3: Pregenerate Pool of names (small). Once we going to reset number we will
//    pick a new name from the pool and return old one. If numbers not enought
//    generate some names additionaly
// 4: Most efficiant and Simple (IMHO). Two letters are random serve us as Key
//    for a map where value for each entry is sequential number. We still going
//    to run out of numbers at some moment. But. It's random (letter are random)
//    so we have 1/26 chances to hit sequenced number, but it will be uniq.
//    with no check - we will fall into infinite loop after 26*26*1000 used names.

// UcLetASCII - array defines start and end of ascii codes
// for uppercase letters used in name generation.
var UcLetASCII = [2]int{65, 90}

// var NumsASCII [2]int = [2]int{49, 58}

// Internal Cache Storage so we can store randomness
var storage = make(map[string]int16)

// Robot structure
type Robot struct {
	name string
}

// New will give us new Robot with factory settings.
func New() *Robot {
	r := new(Robot)
	r.name = randName()
	return r
}

// Name - will return name, if its not set we are resetting
// its name.
func (r Robot) Name() string {
	if r.name == "" {
		r.Reset()
	}
	return r.name
}

// We going only to generate letters in out serial number.
// Numbers going to be sequential for each letter series.
// in this case we not going to spend 2 byts for each existing number
// and no need to iterate thought the list.
// two letter pattern satisfy randomnes condition.
// sequential number makes it easier and relieble to use.
func randName() string {
	rand.Seed(time.Now().UTC().UnixNano())

	name := bytes.NewBuffer([]byte{})
	// Generating First two numbers
	name.WriteByte(byte(UcLetASCII[0] + rand.Intn(UcLetASCII[1]-UcLetASCII[0])))
	name.WriteByte(byte(UcLetASCII[0] + rand.Intn(UcLetASCII[1]-UcLetASCII[0])))

	records, ok := storage[name.String()]
	if ok == true {
		if records == 999 {
			return randName()
		}
		storage[name.String()]++
	} else {
		storage[name.String()] = 0
	}

	return fmt.Sprintf("%s%.3d", name.String(), storage[name.String()])
}

// Reset name triggers randName to generate new name / serial for robot.
func (r *Robot) Reset() {
	r.name = randName()
}
