package tests

import (
	"testing"
)

/**
 * [137] Single Number II
 *
 * Given a non-empty array of integers, every element appears three times except for one, which appears exactly once. Find that single one.
 *
 * Note:
 *
 * Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
 *
 * Example 1:
 *
 *
 * Input: [2,2,3,2]
 * Output: 3
 *
 *
 * Example 2:
 *
 *
 * Input: [0,1,0,1,0,1,99]
 * Output: 99
 *
 */

func TestSingleNumberII(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{2,2,3,2},
			output: 3,
		},
		{
			input:  []int{0,1,0,1,0,1,99},
			output: 99,
		},
	}
	for _, c := range cases {
		x := singleNumberII(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func singleNumberII(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	quickSort(nums, 0, len(nums)-1)
	i := 0
	for i < len(nums) - 1 {
		if nums[i] == nums[i+1] {
			i+=3
			continue
		} else  {
			return nums[i]
		}
	}
	return nums[i]
}

// submission codes end
