package main
import "strconv"

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func fizzBuzz(n int) []string {
	var result = make([]string, n)

	for i := 1; i <= n; i++ {
		if i%3 != 0 && i%5 != 0 {
			result[i-1] = strconv.Itoa(i)
			continue
		}

		if i%3 == 0 {
			result[i-1] += "Fizz"
		}

		if i%5 == 0 {
			result[i-1] += "Buzz"
		}
	}
	return result
}
