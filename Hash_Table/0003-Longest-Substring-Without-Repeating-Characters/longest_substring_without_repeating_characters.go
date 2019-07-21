package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/
func lengthOfLongestSubstring(s string) int {
	var (
		max, cur int                 // max and current counters
		az       = make([]bool, 255) // lookup ascii
	)

	for i := 0; i < len(s); i++ {

		if az[s[i]] {
			for s[cur] != s[i] {
				az[s[cur]] = false
				cur++
			}
			cur++
		} else {
			az[s[i]] = true
		}

		if i-cur+1 > max {
			max = i - cur + 1
		}
	}

	return max
}
