package main
import "strings"
import "fmt"

/*******************************************************************************
  Problem Solution
*******************************************************************************/

//	All we need to do is just keep track of sequence nN or Nn
//	enywhere inside of the word.

func detectCapitalUse(word string) bool {

	for i := 1; i < len(word); i++ {
		FirstUp, SecondLow := int(word[i-1]) < 91, int(word[i]) > 95

		if (!FirstUp && !SecondLow) || (FirstUp && SecondLow && i > 1) {
			return false
		}
	}

	return true
}

// All we need to do is just keep track of sequence nN or Nn
// enywhere inside of the word.

func main() {

	// Thanks to https://www.reddit.com/r/PeopleFuckingDying/
	// for test patterns.
	// WiLD dogs brUTALLy TEAr Arm OFF DEfENsELESS HUMAN
	// https://gfycat.com/gifs/detail/GraveColorfulIcelandgull

	var testString string = "WiLD dogs brUTALLy TEAr Arm OFF DEfENsELESS HUMAN"

	for _, word := range strings.Fields(testString) {
		fmt.Printf("%v (%s)\n", detectCapitalUse(word), word)
	}
}
