package main
/*******************************************************************************
  Problem Solution
*******************************************************************************/
func hIndex(citations []int) int {

	var l = len(citations)

	if l == 0 {
		return 0
	}

	var i, li, ui = 0, 0, l - 1

	for li <= ui {
		i = li + (ui-li)/2

		if citations[i] < l-i {
			li = i + 1
			continue
		}
		ui = i - 1

	}

	return l - li
}
