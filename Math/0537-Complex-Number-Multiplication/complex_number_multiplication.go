package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func complexNumberMultiply(aString string, bString string) string {

	var sToCplx = func(in string) (a int, b int) {
		s := strings.Split(in, "+")
		a, _ = strconv.Atoi(s[0])
		b, _ = strconv.Atoi(s[1][:len(s[1])-1])
		return
	}

	a, ai := sToCplx(aString)
	b, bi := sToCplx(bString)

	return fmt.Sprintf("%d+%di", a*b+ai*bi*-1, a*bi+ai*b)
}
