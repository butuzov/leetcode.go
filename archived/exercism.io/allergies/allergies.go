package allergies

import "fmt"

var mapAllergyScores map[string]uint = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

type Allergy struct {
	Name  string
	Score uint
}

var AllergyList = []Allergy{
	{"eggs", 1},
	{"peanuts", 2},
	{"shellfish", 4},
	{"strawberries", 8},
	{"tomatoes", 16},
	{"chocolate", 32},
	{"pollen", 64},
	{"cats", 128},
}

// Allergies - will return slice of strings of allergies substances
// based on Scrore. Uses AllergicTo.
func Allergies(score uint) []string {
	var allergies []string

	// We just iterating thougth the slice of Allergies and testing
	// each one.
	for _, allergy := range AllergyList {
		if AllergicTo(score, allergy.Name) {
			allergies = append(allergies, allergy.Name)
		}
	}

	return allergies
}

// AllergicTo return state based on substance name and
// allergy score.
func AllergicTo(score uint, substance string) bool {
	for _, allergy := range AllergyList {
		if substance == allergy.Name {
			// We using two opperations here
			// 1. We Working with modulo of score divided by 256.
			//    Guess its fastest way to backt o int8 number.
			// 2. Using AND bits opperator we checking is number a
			//    Power of two
			if (score%256)&allergy.Score > 0 {
				return true
			}
			// No need to iterate more...
			break
		}
	}
	return false
}

func main() {

	fmt.Println(AllergyList)

	Allergies(257)
}
