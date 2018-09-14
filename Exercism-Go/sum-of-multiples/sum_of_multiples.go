package summultiples

import (
	"math"
	"sort"
)

// SumMultiplesNaive Returns sum of multiples (naive vay)
func SumMultiplesNaive(limit int, multiples ...int) int {
	// sum value
	var sum int

	// memorization
	var sumMap map[int]int = make(map[int]int)

	// fail fast
	if len(multiples) == 0 {
		return sum
	}

	for _, number := range multiples {
		for i := number; i < limit; i += number {
			sumMap[i]++
			if sumMap[i] == 1 {
				sum += i
			}
		}
	}

	return sum
}

// SumMultiples do some smart things slowly
func SumMultiples(limit int, multiples ...int) int {
	var sum int

	var ints = make([]int, len(multiples))
	copy(ints, multiples)

	// Sorting
	if sort.IntsAreSorted(ints) == false {
		sort.Ints(ints)
	}
	// fmt.Printf("MultiplesSorted(%#v)\n\n", ints)

	// Making list Uniq
	currentIntsIndex := 0
	for i := range ints {
		if currentIntsIndex == 0 || ints[currentIntsIndex-1] != ints[i] {
			ints[currentIntsIndex] = ints[i]
			currentIntsIndex++
		}
	}
	ints = ints[:currentIntsIndex]
	// fmt.Printf("MultiplesUniq(%#v)\n\n", ints)

	// first attempt to fail
	if len(ints) == 0 {
		return sum
	}

	// Generating Combinations
	var intsCombinations = make([][]int, 0)
	Combinations(&intsCombinations, []int{}, ints)

	for i := range intsCombinations {
		// fmt.Printf("Combinations %#v\n", intsCombinations[i])

		// in this part we going to find lcm for a slice
		lcm := LowestComonMultiple(intsCombinations[i])
		// fmt.Printf("%#v - %#v\n\n", lcm, intsCombinations[i])

		if lcm >= limit {
			continue
		}

		divisors, remainder := (limit / lcm), (limit % lcm)
		if remainder == 0 {
			divisors--
		}

		if divisors == 0 {
			return sum
		}

		sumLocal := int((float64(divisors+1) / 2) * float64(lcm*divisors))
		// fmt.Printf("%#v\n", sumLocal)

		// if its odd add, even subtract
		if len(intsCombinations[i])%2 == 1 {
			sum += sumLocal
		} else {
			sum -= sumLocal
		}
	}

	return sum
}

// Combinations of all posible values
func Combinations(stored *[][]int, initial, values []int) {
	for i := range values {
		var value []int = make([]int, len(initial)+1)
		copy(value, initial)
		value[len(value)-1] = values[i]
		*stored = append(*stored, value)
		Combinations(stored, value, values[i+1:])
	}
}

// GreatestCommonDivider its noraml GreatestCommonDivider algorithme
func GreatestCommonDivider(a int, b int) int {
	var gcm = 1

	if b == 0 {
		gcm = a
	} else if a >= b && b > 0 {
		gcm = GreatestCommonDivider(b, a%b)
	} else if b >= a && a > 0 {
		gcm = GreatestCommonDivider(b, a%b)
	}

	return gcm
}

// LowestComonMultiple its noraml LowestComonMultiple algorithme
func LowestComonMultiple(nums []int) int {
	var lcm int

	for i, num := range nums {
		if i == 0 {
			lcm = num
		} else {
			lcm = int(math.Abs(float64(lcm*num))) / GreatestCommonDivider(lcm, num)
		}
	}

	return lcm
}
