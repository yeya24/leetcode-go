package tests

import (
	"math"
	"testing"
)

/**
 * [414] Third Maximum Number
 *
 * Given a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).
 *
 * Example 1:
 *
 * Input: [3, 2, 1]
 *
 * Output: 1
 *
 * Explanation: The third maximum is 1.
 *
 *
 *
 * Example 2:
 *
 * Input: [1, 2]
 *
 * Output: 2
 *
 * Explanation: The third maximum does not exist, so the maximum (2) is returned instead.
 *
 *
 *
 * Example 3:
 *
 * Input: [2, 2, 3, 1]
 *
 * Output: 1
 *
 * Explanation: Note that the third maximum here means the third maximum distinct number.
 * Both numbers with value 2 are both considered as second maximum.
 *
 *
 */

func TestThirdMaximumNumber(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{3, 2, 1},
			output: 1,
		},
		{
			input:  []int{1, 2},
			output: 2,
		},
		{
			input:  []int{2, 2, 3, 1},
			output: 1,
		},
	}
	for _, c := range cases {
		x := thirdMax(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func thirdMax(nums []int) int {
	first := math.MinInt64
	second := math.MinInt64
	third := math.MinInt64
	for _, v := range nums {
		if v > first {
			third = second
			second = first
			first = v
		} else if v < first && v > second {
			third = second
			second = v
		} else if v < second && v > third {
			third = v
		}
	}
	if third == math.MinInt64 {
		return first
	}
	return third
}

// submission codes end
