package reachingpoints

/*******************************************************************************
  TestTable
*******************************************************************************/

type Case struct {
	sx       int
	sy       int
	tx       int
	ty       int
	expected bool
}

var TestCases = []Case{
	// simple
	{sx: 1, sy: 1, tx: 1, ty: 1, expected: true},
	{sx: 1, sy: 1, tx: 9, ty: 13, expected: true},
	{sx: 2, sy: 4, tx: 5, ty: 9, expected: false},
	{sx: 2, sy: 7, tx: 9, ty: 16, expected: true},
	{sx: 9, sy: 10, tx: 9, ty: 19, expected: true},

	// uneven spliting 70/30
	{sx: 1, sy: 1, tx: 19, ty: 9, expected: true},
	{sx: 1, sy: 1, tx: 109, ty: 9, expected: true},
	{sx: 1, sy: 1, tx: 199, ty: 9, expected: true},
	{sx: 1, sy: 1, tx: 289, ty: 9, expected: true},
	{sx: 1, sy: 1, tx: 379, ty: 9, expected: true},
	{sx: 1, sy: 1, tx: 469, ty: 9, expected: true},
	{sx: 1, sy: 1, tx: 559, ty: 9, expected: true},
	{sx: 1, sy: 1, tx: 649, ty: 9, expected: true},
	{sx: 1, sy: 1, tx: 739, ty: 9, expected: true},

	// random
	{sx: 1, sy: 1, tx: 60, ty: 11, expected: true},
	{sx: 1, sy: 1, tx: 4006, ty: 1077, expected: true},
	{sx: 1, sy: 1, tx: 11573, ty: 26204, expected: true},

	{sx: 1, sy: 1, tx: 1061995, ty: 296566, expected: true},
	{sx: 1, sy: 1, tx: 2859236, ty: 4208893, expected: true},
	{sx: 1, sy: 1, tx: 44302654, ty: 12293587, expected: true},
	{sx: 1, sy: 1, tx: 2923873624, ty: 1850077389, expected: true},
	{sx: 1, sy: 1, tx: 29319664096, ty: 18812802233, expected: true},
}
