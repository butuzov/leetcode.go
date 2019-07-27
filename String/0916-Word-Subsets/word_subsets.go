package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/

type Word struct {
	abc [26]int // Lookup table
}

func newWord(input string) Word {
	var word = Word{}

	for i := 0; i < len(input); i++ {
		word.abc[int(input[i])-'a']++
	}

	return word
}

func (w Word) isSubsetOf(cmp Word) bool {
	for i := 0; i < len(w.abc); i++ {
		if cmp.abc[i]-w.abc[i] < 0 {
			return false
		}
	}
	return true
}

func (w *Word) addMax(n Word) {
	for i := 0; i < len(w.abc); i++ {
		if w.abc[i] < n.abc[i] {
			w.abc[i] = n.abc[i]
		}
	}
}

// Compact all checks into one variable
func Compact(wordlist []string) Word {
	var w = newWord("")
	for i := range wordlist {
		w.addMax(newWord(wordlist[i]))
	}
	return w
}

func wordSubsets(A []string, B []string) []string {
	var words []string

	// Lookup table data
	var cmp = Compact(B)

	for i := range A {
		if cmp.isSubsetOf(newWord(A[i])) {
			words = append(words, A[i])
		}
	}

	return words
}
