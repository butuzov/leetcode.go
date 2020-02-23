package main

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func groupAnagrams(strs []string) [][]string {

	var lookups = make(map[[26]int][]string)
	for i := 0; i < len(strs); i++ {
		var key [26]int

		for chr := 0; chr < len(strs[i]); chr++ {
			key[int8(strs[i][chr]%97)]++
			// key ^= int32(strs[i][chr] % 97)
		}

		if _, ok := lookups[key]; !ok {
			lookups[key] = make([]string, 0, 1)
		}
		lookups[key] = append(lookups[key], strs[i])
	}

	var results = make([][]string, 0, len(lookups))
	for _, v := range lookups {
		results = append(results, v)
	}

	return results
}
