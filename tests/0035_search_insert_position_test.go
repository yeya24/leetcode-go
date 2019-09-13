package tests

import (
	"testing"
)

/**
 * [35] Search Insert Position
 *
 * Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
 *
 * You may assume no duplicates in the array.
 *
 * Example 1:
 *
 *
 * Input: [1,3,5,6], 5
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: [1,3,5,6], 2
 * Output: 1
 *
 *
 * Example 3:
 *
 *
 * Input: [1,3,5,6], 7
 * Output: 4
 *
 *
 * Example 4:
 *
 *
 * Input: [1,3,5,6], 0
 * Output: 0
 *
 *
 */

func TestSearchInsertPosition(t *testing.T) {
	var cases = []struct {
		input  []int
		target int
		output int
	}{
		{
			input:  []int{1, 3, 5, 6},
			target: 5,
			output: 2,
		},
		{
			input:  []int{1, 3, 5, 6},
			target: 2,
			output: 1,
		},
		{
			input:  []int{1, 3, 5, 6},
			target: 7,
			output: 4,
		},
		{
			input:  []int{1, 3, 5, 6},
			target: 0,
			output: 0,
		},
		{
		    input: []int{},
		    target: 1,
		    output: 0,
        },
	}
	for _, c := range cases {
		x := searchInsert(c.input, c.target)
		if x != c.output {
			t.Fail()
		}
	}
}

// submission codes start here

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := i + (j-i)/2
		if target == nums[m] {
			return m
		} else if nums[m] > target {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	return i
}

// submission codes end
