package main

import "sort"

/*******************************************************************************
  Problem Solution
*******************************************************************************/

//
func lexicalOrder(n int) []int {

	var base = func(n int) int {
		var pow = 10
		for n%pow != n {
			pow *= 10
		}
		return pow / 10
	}

	var result = make([]int, n)
	var lookup = make([]int, n)
	var maxnum = base(n)

	for i := 1; i <= n; i++ {
		result[i-1] = i
		lookup[i-1] = i * maxnum / base(i)
	}

	sort.Slice(result, func(i, j int) bool {
		if lookup[result[i]-1] == lookup[result[j]-1] {
			return result[i] < result[j]
		}
		return lookup[result[i]-1] < lookup[result[j]-1]
	})

	return result
}

// lexicalOrderDFS by https://github.com/aQuaYi
func lexicalOrderDFS(max int) []int {
	res := make([]int, 0, max)

	var dfs func(int)
	dfs = func(x int) {
		limit := (x + 10) / 10 * 10
		for x <= max && x < limit {
			res = append(res, x)
			if x*10 <= max {
				dfs(x * 10)
			}
			x++
		}
	}
	dfs(1)
	return res
}
