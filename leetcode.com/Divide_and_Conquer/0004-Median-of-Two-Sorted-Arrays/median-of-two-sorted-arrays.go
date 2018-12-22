package medianoftwosortedarrays

import "sort"

// Naive implementations (just for testing)
func findMedianSortedArraysNaive(nums1 []int, nums2 []int) float64 {
	medianas := medianaAssumptions(len(nums1), len(nums2))

	nums1 = append(nums1, nums2...)
	sort.Sort(sort.IntSlice(nums1))
	var mediana int

	for _, i := range medianas {
		mediana += nums1[i]
	}

	return float64(mediana) / float64(len(medianas))
}

// This is top performant solution on leetcode for Go.

func findMedianSortedArraysTop(nums1 []int, nums2 []int) float64 {
	lenlist := len(nums1) + len(nums2)
	//fmt.Println(lenlist)
	nums := make([]int, lenlist)
	nums1_i, nums2_i := 0, 0
	for i := 0; i < lenlist; i++ {
		if len(nums1) == 0 {
			copy(nums, nums2)
			break
		} else if len(nums2) == 0 {
			copy(nums, nums1)
			break
		}

		if nums1[nums1_i] <= nums2[nums2_i] {
			nums[i] = nums1[nums1_i]
			if nums1_i == len(nums1)-1 {
				copy(nums[i+1:], nums2[nums2_i:])
				break
			}

			nums1_i += 1

		} else {
			nums[i] = nums2[nums2_i]
			if nums2_i == len(nums2)-1 {
				copy(nums[i+1:], nums1[nums1_i:])
				break
			}

			nums2_i += 1
		}
	}
	//fmt.Println(nums)
	result := 0.0
	if lenlist%2 == 0 {
		i := lenlist / 2
		result = float64(nums[i-1]+nums[i]) / 2.0
	} else {
		i := lenlist / 2
		result = float64(nums[i])
	}

	return result
}

// This implementation isn't really nice looking, bt its easy to understand.
// so we trying to get alement by particular index, without merging element.
// we picking middle value of smaller array and trying to guess where its
// going to be in larger array. if the index of new element is bigger than
// index we looking for - we reducing both arrays, and decrising index.
// so it looks like more dived and conquere.

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	medianas := medianaAssumptions(len(nums1), len(nums2))
	var comlink = make(chan int, len(medianas))
	var mediana int

	getIndex := func(c chan int, i int, nums1 []int, nums2 []int) {
		c <- getVirtualSliceIndex(i, nums1, nums2)
	}

	for index := range medianas {
		go getIndex(comlink, medianas[index], nums1, nums2)
	}

	for i := 0; i < len(medianas); i++ {
		mediana += <-comlink
	}
	close(comlink)

	return float64(mediana) / float64(len(medianas))
}

// medianaAssumptions giving us a slice with a numbers (array indexes)
// of elements that suppose to make median (and ones we going to look for)
func medianaAssumptions(a, b int) []int {
	t := a + b
	medianas := []int{}
	if t%2 == 0 {
		medianas = append(medianas, (t/2)-1)
	}
	medianas = append(medianas, (t / 2))
	return medianas
}

// binarySearch encapsulate native go implementation of binary search.
func binarySearch(num int, nums []int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] >= num
	})
}

// Will return element at some index in virtual slice created by a and b
// without creation
func getVirtualSliceIndex(index int, a, b []int) int {

	aLen, bLen := len(a), len(b)

	if aLen > bLen {
		b, a = a, b
		aLen, bLen = bLen, aLen
	}

	if aLen == 0 {
		return b[index]
	}

	if bLen == 0 {
		return a[index]
	}

	if aLen == 1 {
		if a[0] > b[bLen-1] {
			if index > bLen-1 {
				return a[0]
			}
			return b[index]
		}

		if a[0] < b[0] {
			if index > aLen-1 {
				return b[index-1]
			}
			return a[0]
		}
	}

	aIndexMuddle := aLen / 2
	bIndexAssumed := binarySearch(a[aIndexMuddle], b)
	indexAssumed := aIndexMuddle + bIndexAssumed

	if index == indexAssumed {
		return a[aIndexMuddle]
	}

	if index < indexAssumed {
		a = a[:aIndexMuddle]
		b = b[:bIndexAssumed]
	}

	if index > indexAssumed {
		a = a[aIndexMuddle:]
		b = b[bIndexAssumed:]
		index -= indexAssumed

		// Post reduce correction
		if len(a) == 1 && a[0] == b[0] {
			a = a[1:]
			index--
		}
	}

	return getVirtualSliceIndex(index, a, b)
}
