package tests

import (
	"math"
	"testing"
)

/**
 * [628] Maximum Product of Three Numbers
 *
 * Given an integer array, find three numbers whose product is maximum and output the maximum product.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3]
 * Output: 6
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4]
 * Output: 24
 *
 *
 *
 *
 * Note:
 *
 *
 * 	The length of the given array will be in range [3,10^4] and all elements are in the range [-1000, 1000].
 * 	Multiplication of any three numbers in the input won't exceed the range of 32-bit signed integer.
 *
 *
 *
 *
 */

func TestMaximumProductofThreeNumbers(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 2, 3},
			output: 6,
		},
		{
			input:  []int{1, 2, 3, 4},
			output: 24,
		},
		{
			input:  []int{-100, -1, 3, 4, 5},
			output: 500,
		},
		{
			input:  []int{0, 0, 0, 4},
			output: 0,
		},
	}
	for _, c := range cases {
		x := maximumProduct(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func maximumProduct(nums []int) int {
	max, max2, max3, min, min2 := -1001, -1001, -1001, 1001, 1001
	l := len(nums)
	if l == 3 {
		return nums[0] * nums[1] * nums[2]
	}
	for _, n := range nums {
		if n >= max {
			max3 = max2
			max2 = max
			max = n
		} else if n >= max2 {
			max3 = max2
			max2 = n
		} else if n >= max3 {
			max3 = n
		}
		if n <= min {
			min2 = min
			min = n
		} else if n <= min2 {
			min2 = n
		}
	}
	return int(math.Max(float64(max*max2*max3), float64(min*min2*max)))
}

// submission codes end
