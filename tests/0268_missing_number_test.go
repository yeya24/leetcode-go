package tests

import (
	"testing"
)

/**
 * [268] Missing Number
 *
 * Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.
 *
 * Example 1:
 *
 *
 * Input: [3,0,1]
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: [9,6,4,2,3,5,7,0,1]
 * Output: 8
 *
 *
 * Note:
 * Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?
 */

func TestMissingNumber(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{3, 0, 1},
			output: 2,
		},
		{
			input:  []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			output: 8,
		},
	}
	for _, c := range cases {
		x := missingNumber(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func missingNumber(nums []int) int {
	l := len(nums)
	for i, v := range nums {
		l = l ^ i ^ v
	}
	return l
}

// submission codes end
