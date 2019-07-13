package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/
const ABC_LENGTH = 26

func convertToTitle(n int) string {
	if n <= 0 {
		return ""
	}
	return convertToTitle((n-1)/ABC_LENGTH) + string('A'+(n-1)%ABC_LENGTH)
}
