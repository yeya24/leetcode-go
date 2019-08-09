package tests

import (
	"testing"
)

/**
 * [485] Max Consecutive Ones
 *
 * Given a binary array, find the maximum number of consecutive 1s in this array.
 *
 * Example 1:
 *
 * Input: [1,1,0,1,1,1]
 * Output: 3
 * Explanation: The first two digits or the last three digits are consecutive 1s.
 *     The maximum number of consecutive 1s is 3.
 *
 *
 *
 * Note:
 *
 * The input array will only contain 0 and 1.
 * The length of input array is a positive integer and will not exceed 10,000
 *
 *
 */

func TestMaxConsecutiveOnes(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 1, 0, 1, 1, 1},
			output: 3,
		},
	}
	for _, c := range cases {
		x := findMaxConsecutiveOnes(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func findMaxConsecutiveOnes(nums []int) int {
	c := 0
	a := 0
	for _, v := range nums {
		if v == 1 {
			a++
		} else {
			a = 0
		}
		if a > c {
			c = a
		}
	}
	return c
}

// submission codes end
