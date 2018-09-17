package armstrong

import (
	"fmt"
	"math"
)

func main() {
	if IsNumber(9) {
		fmt.Println("ok")
	}
}

// IsNumber determine is its a Armstring Number
func IsNumber(num int) bool {

	var (
		nums     []int
		original int
	)
	// preserving original num variable
	original = num

	// fail if its negative
	if num < 0 {
		return false
	}

	// While loop to determine dividers
	for {
		nums = append(nums, num%10)
		num = (num - nums[len(nums)-1]) / 10

		// stop condition.
		if num == 0 {
			break
		}
	}

	// summing slice of ints.
	var sum int
	for _, i := range nums {
		sum += int(math.Pow(float64(i), float64(len(nums))))
	}

	// is it Armstrong number?
	return sum == original
}
