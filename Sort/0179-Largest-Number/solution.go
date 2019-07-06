package main
import (
	"math"
	"sort"
	"strconv"
)

/*******************************************************************************
  Problem Solution
*******************************************************************************/

type numStrings struct {
	keys    []int
	strings map[int]string
}

func (ns numStrings) Less(i, j int) bool {
	k1, k2 := ns.keys[i], ns.keys[j]

	var k1Cmp = k2 + k1*int(math.Pow10(len(ns.strings[k2])))
	var k2Cmp = k1 + k2*int(math.Pow10(len(ns.strings[k1])))

	return k1Cmp > k2Cmp
}

func (ns numStrings) Swap(i, j int) {
	ns.keys[i], ns.keys[j] = ns.keys[j], ns.keys[i]
}

func (ns numStrings) Len() int {
	return len(ns.keys)
}

func largestNumber(nums []int) string {

	ns := numStrings{
		keys:    make([]int, len(nums)),
		strings: make(map[int]string),
	}

	for i, num := range nums {
		ns.keys[i] = num
		ns.strings[num] = strconv.Itoa(num)
	}

	sort.Sort(numStrings(ns))

	//	fmt.Printf("%+v\n", ns)

	var result string
	for _, v := range ns.keys {
		if v == 0 && result == "0" {
			continue
		}
		result += ns.strings[v]
	}

	return result
}
