// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear pretty simple fucntion to determine if year is leap or not.
func IsLeapYear(year int) bool {

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}

	return false
}
