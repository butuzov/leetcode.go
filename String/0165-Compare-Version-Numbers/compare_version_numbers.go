package main

import (
	"strconv"
	"strings"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func removeLeadZeros(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] != '0' {
			return s[i:]
		}
	}

	return string('0')
}

func parseVersion(version string, v *[]int) {
	for i, nString := range strings.Split(version, ".") {
		tmpInt, _ := strconv.Atoi(removeLeadZeros(nString))
		(*v)[i] = tmpInt
	}
}

func compareVersion(version1 string, version2 string) int {

	var biggestVersionLen = max(strings.Count(version1, "."), strings.Count(version2, "."))
	biggestVersionLen++

	var versionOneInts = make([]int, biggestVersionLen)
	parseVersion(version1, &versionOneInts)

	var versionTwoInts = make([]int, biggestVersionLen)
	parseVersion(version2, &versionTwoInts)

	for len(versionOneInts) > 0 {
		if versionOneInts[0] < versionTwoInts[0] {
			return -1
		}

		if versionOneInts[0] > versionTwoInts[0] {
			return 1
		}
		versionOneInts = versionOneInts[1:]
		versionTwoInts = versionTwoInts[1:]
	}

	return 0
}
