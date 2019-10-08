package tests

import (
	"math"
	"testing"
)

/**
 * [209] Minimum Size Subarray Sum
 *
 * Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum >= s. If there isn't one, return 0 instead.
 *
 * Example:
 *
 *
 * Input: s = 7, nums = [2,3,1,2,4,3]
 * Output: 2
 * Explanation: the subarray [4,3] has the minimal length under the problem constraint.
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n).
 *
 */

func TestMinimumSizeSubarraySum(t *testing.T) {
	var cases = []struct {
		input  []int
		s      int
		output int
	}{
		{
			input:  []int{2, 3, 1, 2, 4, 3},
			s:      7,
			output: 2,
		},
		{
			s:      213,
			input:  []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12},
			output: 8,
		},
	}
	for _, c := range cases {
		x := minSubArrayLen(c.s, c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func minSubArrayLen(s int, nums []int) int {
	left := 0
	sum := 0
	c := math.MaxInt64
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= s {
			c = min(c, right+1-left)
			sum -= nums[left]
			left++
		}
	}
	if c == math.MaxInt64 {
		return 0
	}
	return c
}

// submission codes end
