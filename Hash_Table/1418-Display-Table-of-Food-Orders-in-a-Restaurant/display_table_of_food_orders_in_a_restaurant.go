package main

import (
	"sort"
	"strconv"
)

// ******************************************
// https://github.com/butuzov/leetcode.go
// ******************************************

func displayTable(in [][]string) [][]string {

	var (
		mapPlates = make(map[string]int) // <- lookup map for plate -> index of the columns
		valPlates = make([]string, 0, 4) // <- values of plates used for sorting: e.g. Egg benedict
	)
	var (
		mapTables = make(map[int]int) // <- lookup map for table -> index of the row
		valTables = make([]int, 0, 4) // <- values of tables used for sorting: e.g. 10, 20, 1
	)

	var orders = make([][]int, len(in))

	for i := range in {
		orders[i] = make([]int, len(in))

		table, _ := strconv.Atoi(in[i][1])
		plate := in[i][2]

		var posPlate, posTable int
		var ok bool

		if posPlate, ok = mapPlates[plate]; !ok {
			posPlate = len(valPlates)
			mapPlates[plate] = posPlate
			valPlates = append(valPlates, plate)
		}

		if posTable, ok = mapTables[table]; !ok {
			posTable = len(valTables)
			mapTables[table] = posTable
			valTables = append(valTables, table)
		}

		orders[posPlate][posTable]++
	}

	var table = make([][]string, len(valTables)+1)

	// Header
	table[0] = make([]string, len(valPlates)+1)
	table[0][0] = "Table"
	sort.Strings(valPlates)
	for i := 1; i <= len(valPlates); i++ {
		table[0][i] = valPlates[i-1]
	}

	sort.Ints(valTables)
	for col := len(valTables); col >= 1; col-- {
		table[col] = make([]string, len(valPlates)+1)
		table[col][0] = strconv.Itoa(valTables[col-1])

		for i := 1; i <= len(valPlates); i++ {
			plateKey := mapPlates[valPlates[i-1]]
			tableKey := mapTables[valTables[col-1]]
			table[col][i] = strconv.Itoa(orders[plateKey][tableKey])
		}
	}

	return table
}
