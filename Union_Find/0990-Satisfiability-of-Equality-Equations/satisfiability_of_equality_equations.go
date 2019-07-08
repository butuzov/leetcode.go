package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func equationsPossible(equations []string) bool {
	var lookups = make([]int, 26)
	for i := range lookups {
		lookups[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if lookups[x] != x {
			lookups[x] = find(lookups[x])
		}
		return lookups[x]
	}

	union := func(a, b int) {
		lookups[find(a)] = find(b)
	}

	for _, e := range equations {
		if e[1] != '=' {
			continue
		}
		union(int(e[0]-'a'), int(e[3]-'a'))
	}

	for _, e := range equations {
		if e[1] == '=' {
			continue
		}

		if find(int(e[0]-'a')) == find(int(e[3]-'a')) {
			return false
		}
	}

	return true
}
