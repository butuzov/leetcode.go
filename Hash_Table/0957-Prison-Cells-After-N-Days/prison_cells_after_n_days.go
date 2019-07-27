package main

/*******************************************************************************
  Problem Solution
*******************************************************************************/

// prisonAfterNDaysUnknownAuthor - top memoty on leetcode. Author unknown.
func prisonAfterNDaysUnknownAuthor(cells []int, N int) []int {
	var data = make([]int, len(cells))
	copy(data, cells)

	for n := 0; n < (N-1)%14+1; n++ {
		tmpCells := append(data[:0:0], data...)
		for i := 1; i < 7; i++ {
			data[i] = tmpCells[i-1] ^ tmpCells[i+1] ^ 1
		}
		data[0] = 0
		data[7] = 0
	}

	return data
}

// prisonAfterNDaysNaive - is Naive solution
func prisonAfterNDaysNaive(cells []int, n int) []int {

	var data = make([]int, len(cells))
	copy(data, cells)
	var row = make([]int, len(data))

	var vars = map[bool]int{
		true:  1,
		false: 0,
	}

	for i := 0; i < (n-1)%14+1; i++ {
		for p := 1; p < len(data)-1; p++ {
			row[p] = vars[(data[p-1] == data[p+1])]
		}
		copy(data, row)
	}
	return row
}

// prisonAfterNDays - default implementation of the problem solution
func prisonAfterNDays(cells []int, n int) []int {

	var data = make([]int, len(cells))
	copy(data, cells)

	var num = ConvertToInt(data)

	// fmt.Println("      0 ", intToBitArray(num))
	// fmt.Println(strings.Repeat("=", 40))

	for i := 0; i < (n-1)%14+1; i++ {
		// fmt.Println(strings.Repeat("=", 40))
		// fmt.Printf("Shift Left  %8[1]d - %08[1]b\n", uint8(num<<1))
		// fmt.Printf("Shift Right %8[1]d - %08[1]b\n", uint8(num>>1))
		var tmp = (uint8(num>>1) ^ uint8(num<<1))

		// fmt.Printf("s^255       %8[1]d - %08[1]b\n", tmp)
		tmp = tmp ^ 255
		// fmt.Printf("s^255       %8[1]d - %08[1]b\n", tmp)

		// naive left and right clear
		tmp = tmp >> 1
		tmp = tmp << 1
		// fmt.Printf("Shifted     %8[1]d - %08[1]b\n", tmp)

		tmp = tmp << 1
		tmp = tmp >> 1
		// fmt.Printf("Shifted     %8[1]d - %08[1]b\n", tmp)

		num = int(tmp)

		// fmt.Printf("%8[1]d %+v\n", i+1, intToBitArray(num))
		// if i%14 == 13 {
		// 	fmt.Println(strings.Repeat("=", 40))
		// }

	}
	return intToBitArray(num)
}

func ConvertToInt(arr []int) int {
	ret := 0
	factor := 1
	for i := len(arr) - 1; i >= 0; i-- {
		ret += arr[i] * factor
		factor *= 2
	}

	return ret
}

func intToBitArray(val int) []int {
	var result = make([]int, 8)
	var step = len(result) - 1
	for step > 0 {
		result[step], val = val%2, val/2
		step--
	}
	return result
}
